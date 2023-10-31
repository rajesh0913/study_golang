package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name       string
	Password   string
	Phone      string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email      string `valid:"email"`
	Identity   string
	ClientIP   string
	ClientPort string
	Salt       string // 加密字段

	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LoginoutTime  *time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}

	return data
}

// 查询用户

func FindUserByNameAndPwd(name string, password string) *UserBasic {
	user := &UserBasic{}
	utils.DB.Where("name = ? and password = ?", name, password).First(&user)

	// token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	// fmt.Printf("\n\n%s", temp)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func FindUserByName(name string) *UserBasic {
	user := &UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) *UserBasic {
	user := &UserBasic{}
	utils.DB.Where("phone = ?", phone).First(&user)
	return user
}

func FindUserByEmail(email string) *UserBasic {
	user := &UserBasic{}
	utils.DB.Where("name = ?", email).First(&user)
	return user
}

func CreateUser(user UserBasic) *gorm.DB {

	return utils.DB.Create(&user)

}

func DeletUser(user UserBasic) *gorm.DB {

	return utils.DB.Delete(&user)

}
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password, Phone: user.Phone, Email: user.Email})
}
