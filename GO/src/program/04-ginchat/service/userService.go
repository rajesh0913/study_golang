package service

import (
	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"
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
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	// data := make([]*models.UserBasic, 10)
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	rePassord := c.Query("repassword")
	if password != rePassord {

		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.Password = password
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功！",
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
		"message": "用户删除成功！",
	})

}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "用户信息修改成功！",
	})

}
