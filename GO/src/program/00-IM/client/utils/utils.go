package utils

import (
	"GO/src/program/00-IM/pkg/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 创建结构体工厂模式，使用本包函数,将方法关联到结构体中
type Transfer struct {
	Conn net.Conn   // 传输需要的通讯接口
	Buf  [8096]byte // 数据切片,传输缓冲区
}

// 发送数据包
func (t *Transfer) WritePkg(data []byte) (err error) {
	// 1.先发送数据长度给对方
	// 获取长度
	var pkgLen uint32 = uint32(len(data))
	// 将长度进行字符化
	binary.BigEndian.PutUint32((*t).Buf[:4], pkgLen)
	//发送数据长度信息
	n, err := (*t).Conn.Write((*t).Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buf) fail err: ", err)
		return
	}

	// 2.发送数据本身
	n, err = (*t).Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.write(data) fail err: ", err)
		return
	}

	return
}

// 处理数据包
func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	// 读取通讯信息
	// fmt.Println("等待读取客户端数据...")
	// conn.Read任意一方关闭conn就不会阻塞，必须两边都打开才会阻塞等待读取
	n, err := (*t).Conn.Read((*t).Buf[:4])
	if n != 4 || err != nil {
		// err = errors.New("read pkg header error")
		return
	}
	// 根据读到的切片获取信息长度
	var pkgLen uint32 = binary.BigEndian.Uint32((*t).Buf[:4])
	n, err = (*t).Conn.Read((*t).Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// err = errors.New("read pkg body error")
		return
	}
	// 将package反序列化
	err = json.Unmarshal((*t).Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("json.unmashal(buf[:pkgLen]) err:%v\n", err)
		return
	}

	return
}
