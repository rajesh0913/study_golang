package model

import (
	"GO/src/program/00-IM/pkg/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
