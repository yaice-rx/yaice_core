package model

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/satori/go.uuid"
	"github.com/xtaci/kcp-go"
	"time"
)

//个人连接
type Conn struct {
	ConnectInterface
	guid              string
	session           *kcp.UDPSession
	updateConnectTime int64
}

type ConnectInterface interface {
	ReadMsg() ([]byte, error)
	WriteMsg(protoMsg proto.Message)
	Close()
	Destroy()
}

//初始化用户连接信息
func InitConn(conn *kcp.UDPSession) *Conn {
	return &Conn{
		guid:              uuid.Must(uuid.NewV4()).String(),
		session:           conn,
		updateConnectTime: time.Now().Unix(),
	}
}

func (conn *Conn) GetPlayerGuid() string {
	return conn.guid
}

//发送数据
func (conn *Conn) WriteMsg(data []byte) {
	if conn != nil {
		_, err := conn.session.Write(data)
		if err != nil {
			fmt.Println("send msg error ", err.Error())
		}
	} else {
		fmt.Println("kcp connect is nil ", time.Now().String())
	}
}
