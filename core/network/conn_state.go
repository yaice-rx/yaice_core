package network

import (
	"YaIce/core/model"
	"github.com/xtaci/kcp-go"
	"sync"
)

type NetworkConnectList struct {
	MutexConns         sync.Mutex
	MaxConnect         int //最大连接数据
	KcpListen          *kcp.Listener
	ConnectList        map[*kcp.UDPSession]*model.PlayerConn //uid->连接Conn
	ChanMsgReviceQueue chan *model.MsgQueue                  //消息队列
	ChanMshSendQueue   chan *model.MsgQueue                  //发送消息队列
}

var kcpConnsPtr *NetworkConnectList

func Init(maxConn int) {
	kcpConnsPtr = &NetworkConnectList{
		MaxConnect:         maxConn,
		ConnectList:        make(map[*kcp.UDPSession]*model.PlayerConn),
		ChanMsgReviceQueue: make(chan *model.MsgQueue, 10),
	}
}
