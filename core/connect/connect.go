package connect

import (
	"YaIce/core/common"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/satori/go.uuid"
	"github.com/xtaci/kcp-go"
	"time"
)

//个人连接
type PlayerConn struct {
	ConnectInterface
	guid   	string
	session *kcp.UDPSession
	updateConnectTime int64
}

type ConnectInterface interface {
	ReadMsg() ([]byte, error)
	WriteMsg(protoMsg proto.Message)
	Close()
	Destroy()
}

func (conn *PlayerConn)GetPlayerGuid()string{
	return conn.guid
}

//初始化用户连接信息
func InitPlayerConn(conn *kcp.UDPSession)*PlayerConn{
	return &PlayerConn{
		guid:uuid.Must(uuid.NewV4()).String(),
		session:conn,
		updateConnectTime:time.Now().Unix(),
	}
}

//发送数据
func (conn *PlayerConn)WriteMsg(protoMsg proto.Message)  {
	protoNumber := common.ProtocalNumber(common.GetProtoName(protoMsg))
	data,_ := proto.Marshal(protoMsg)
	if conn != nil {
		content := common.IntToBytes(protoNumber)
		content = append(content,data...)
		_,err := conn.session.Write(content)
		if err != nil{
			fmt.Println("send msg error ", err.Error())
		}
	} else {
		fmt.Println("kcp connect is nil ",time.Now().String())
	}
}