package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetUserList
// @Summary 获取用户信息
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	// data := make([]*models.UserBasic, 10)
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param identity query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	// data := make([]*models.UserBasic, 10)
	user := models.UserBasic{}
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("Identity")

	if user.Name == "" || password == "" || repassword == "" {
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "用户名或密码不能为空！",
			"data":    user,
		})
		return
	}

	findUser := models.FindUserByName(user.Name)
	if findUser.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1, // fail
			"message": "用户名已注册！",
			"data":    user,
		})
		return
	}

	if password != repassword {

		c.JSON(200, gin.H{
			"code":        -1, // fail
			"message":     "两次密码不一致",
			"password":    password,
			"repassword:": repassword,
			"data":        user,
		})
		return
	}
	// user.Password = password
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	user.Password = utils.MakePassword(password, salt)

	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0, // success
		"message": "新增用户成功！ ",
		"data":    user,
	})

}

// DeletUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeletUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeletUser(user)
	c.JSON(200, gin.H{
		"code":    0, // success
		"message": "用户删除成功 ",
		"data":    user,
	})

}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	fmt.Println("update: ", user)
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println("UpdateUser govalidator validate fail err: ", err)
		c.JSON(200, gin.H{
			"code":    -1, // success
			"message": "修改参数不匹配！ ",
			"data":    user,
		})
	} else {

		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0, // success
			"message": "用户信息修改成功！ ",
			"data":    user,
		})
	}

}

// FindUserByNameAndPwd
// @Summary 用户登录
// @Tags 用户模块
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {

	// name := c.PostForm("name")
	// password := c.PostForm("password")
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	fmt.Println(name, password)

	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, // fail
			"message": "用户不存在 ",
		})
		return
	}
	if ok := utils.ValidPassword(password, user.Salt, user.Password); !ok {
		c.JSON(200, gin.H{
			"code":    -1, // dail
			"message": "密码不正确  ",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data := models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"code":    0, // success
		"message": "登录成功 ",
		"data":    data,
	})
}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// sendMsg
// @Summary 发送消息
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/sendMsg [get]
func SendMsg(c *gin.Context) {
	// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("SendMsg upgrade fail err: ", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println("SendMsg wsclose fail err: ", err)
			return
		}
	}(ws)
	MsgHandler(ws, c)
}

// 通过redis转发消息
func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println("msgHandler fail err: ", err)
			return
		}
		// fmt.Println("发送消息： ", msg)
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		// 发送客户端订阅通道的信息到服务器
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			// fmt.Println("ws.writemessage fail err: ", err)
			log.Fatalln(err)
			return
		}
	}

}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func SearchFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := models.SearchFriend(uint(id))

	// c.JSON(200, gin.H{
	// 	"code":    0,
	// 	"message": "查询好友列表成功！",
	// 	"data":    users,
	// })
	utils.RespOKList(c.Writer, users, len(users))
}

func AddFriend(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	targetName := c.Request.FormValue("targetName")
	// targetId, _ := strconv.Atoi(c.Request.FormValue("targetId"))
	code, msg := models.AddFriend(uint(userId), targetName)
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

func CreateCommunity(c *gin.Context) {
	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
	name := c.Request.FormValue("name")
	community := models.Community{
		OwnerId: uint(ownerId),
		Name:    name,
	}

	code, msg := models.CreateCommunity(community)
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

func LoadCommunity(c *gin.Context) {
	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))

	data, msg := models.LoadCommunity(uint(ownerId))
	if len(data) != 0 {
		utils.RespList(c.Writer, 0, data, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}
