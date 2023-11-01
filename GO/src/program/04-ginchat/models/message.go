package models

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

/*
SendMsg->node.dataqueue,node.conn:websocket->sendProc(chat:go sendProc)

go recvProc(node.conn:websocket):read->dispatch-SendMsg(msg.targetId) + broadMsg?->udpSendChan<-data
->udpSendProc
为什么做消息转发的时候先走了websocket，又还要走broadMsg->udpSendProc
udp的init似乎没有用到
--猜想：做局域网的判断？如果是在局域网通信消息就直接走udp？所以这里的消息转发其实可以把broadMsg注释掉不走udp？


redis的subscribe和publish转发消息的功能什么时候、哪里会用到
--猜想：心跳检测？用于系统和客户端直接通信？

为什么有的地方要用websocket，有的又要有udp，有的又是http的连接
--猜想：udp是局域网做通信，http和websocket不懂
--在对注册、建群、加好友请求做resp回复的时候用的是http短连接
--跳转到聊天界面是建立长链接所以用websocket？发送消息、接收消息用的是websocket长链接？为何

SendMsg、SendUserMsg其实和redis的订阅和转发没有关系？

*/

type Message struct {
	gorm.Model
	FromId   int64
	TargetId int64
	Type     int // 发送类型 1私聊2群聊3广播
	Media    int //  消息类型 1文字2图片3表情包4音频
	Content  string
	Pic      string
	Url      string
	Desc     string
	Amount   int
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

func Chat(w http.ResponseWriter, request *http.Request) {

	//1.校验token
	query := request.URL.Query()
	id := query.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)
	// msgType := query.Get("type")
	// targetId := query.Get("targetId")
	// context := query.Get("context")
	// 校验请求合法，获取websocket链接
	isValid := true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValid
		},
	}).Upgrade(w, request, nil)
	if err != nil {
		fmt.Println("Char CheckOrigin fail err: ", err)
		return
	}
	// 2.获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 3.用户关系
	// 4.绑定userId和node，加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	// 5. send
	go sendProc(node)
	// 6.recieve
	go recvProc(node)
	// welcome
	fmt.Println("sendMsg >>>>>>>>> userId:", userId)
	sendMsg(userId, []byte("welcome..。"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("[ws] sendProc >>>>>>>>> msg:", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("senProc node.WriteMessage fail err: ", err)
				return
			}

		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
		}
		//心跳检测 msg.Media == -1 || msg.Type == 3
		// if msg.Type == 3 {
		// 	currentTime := uint64(time.Now().Unix())
		// 	node.Heartbeat(currentTime)
		// } else {
		dispatch(data)
		broadMsg(data) //todo 将消息广播到局域网
		fmt.Println("[ws] recvProc <<<<< ", string(data))
		// }

	}
}

var udpSendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpSendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
	fmt.Println("init goroutine...")
}

// 通过udp完成内部数据传输
func udpSendProc() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: 3001,
	})
	if err != nil {
		fmt.Println("udpConn dial fail err: ", err)
		return
	}
	defer udpConn.Close()
	for {
		select {
		case data := <-udpSendChan:
			fmt.Println("udpSenProc >>>>>>>>> data:", string(data))
			_, err := udpConn.Write(data)
			if err != nil {
				fmt.Println("udpSendProc conn Write fail err: ", err)
				return
			}
		}
	}
}

// 完成udp数据接收协程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3001,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("udpRecvProc  data :", string(buf[0:n]))
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	// msg.CreateTime = uint64(time.Now().Unix())
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: //私信
		fmt.Println("dispatch  data :", string(data))
		sendMsg(msg.TargetId, data)
		// case 2: //群发
		// sendGroupMsg(msg.TargetId, data) //发送的群ID ，消息内容
		// case 4: // 心跳
		// 	node.Heartbeat()
		//case 4:
		//
	}
}
func sendMsg(targetId int64, msg []byte) {
	fmt.Println("sendMsg >>> userId: ", targetId, " msg: ", string(msg))
	rwLocker.RLock()
	node, ok := clientMap[targetId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
