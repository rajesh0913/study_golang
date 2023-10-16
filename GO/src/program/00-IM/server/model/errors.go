package model

import "errors"

// 根据业务逻辑需要，自定义error
var (
	ErrorUserNotExists = errors.New("用户不存在")
	ErrorUserExists    = errors.New("用户已经存在")
	ErrorUserPwd       = errors.New("用户密码错误")
)
