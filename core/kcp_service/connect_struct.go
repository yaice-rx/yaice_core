package kcp_service

import (
	"YaIce/core/model"
	"github.com/xtaci/kcp-go"
	"sync"
)

type KcpServiceConnectList struct {
	MutexConns   sync.Mutex
	MaxConnect   int                                   //最大连接数据
	ConnectList  map[*kcp.UDPSession]*model.PlayerConn // uid->连接Conn
	ChanMsgQueue chan *MsgQueue                        //消息队列
}

var KcpConnPtr *KcpServiceConnectList

func InitKcpServiceConn(maxConn int) {
	KcpConnPtr = &KcpServiceConnectList{
		MaxConnect:   maxConn,
		ConnectList:  make(map[*kcp.UDPSession]*model.PlayerConn),
		ChanMsgQueue: make(chan *MsgQueue, 2048),
	}
}
