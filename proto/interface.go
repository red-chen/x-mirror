package proto

import "github.com/red-chen/x-mirror/conn"

type Protocol interface {
	GetName() string
	WrapConn(conn.Conn, interface{}) conn.Conn
}
