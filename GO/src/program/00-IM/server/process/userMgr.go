package process

import "fmt"

// 全局变量，有且只有一个
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 初始化
func InitUserMgr() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成在线用户添加
func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	(*um).onlineUsers[(*up).UserId] = up
}

// 完成在线用户删除
func (um *UserMgr) DeleteOnlineUser(userId int) {
	delete((*um).onlineUsers, userId)
}

// 返回当前所有在线用户
func (um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return (*um).onlineUsers
}

// 根据id返回对应值

func (um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := (*um).onlineUsers[userId]
	// 用户不在线
	if !ok {
		err = fmt.Errorf("用户%d不在线", userId)
		return
	}
	return

}
