package network

import (
	"YaIce/core/model"
	"github.com/xtaci/kcp-go"
	"sync"
)

var MaxCount int = 5000

type NetworkConnectList struct {
	MutexConns         sync.Mutex
	MaxConnect         int //最大连接数据
	KcpListen          *kcp.Listener
	ConnectList        map[*kcp.UDPSession]*model.Conn //uid->连接Conn
	ChanMsgReviceQueue chan *model.MsgQueue            //消息队列
	ChanMsgSendQueue   chan *model.MsgQueue            //发送消息队列
}

var kcpConnsPtr *NetworkConnectList

func Init() {
	kcpConnsPtr = &NetworkConnectList{
		MaxConnect:         MaxCount,
		ConnectList:        make(map[*kcp.UDPSession]*model.Conn),
		ChanMsgReviceQueue: make(chan *model.MsgQueue, 10),
		ChanMsgSendQueue:   make(chan *model.MsgQueue, 10),
	}
}
