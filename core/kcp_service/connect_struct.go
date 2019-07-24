package kcp_service

import (
	"YaIce/core/model"
	"github.com/xtaci/kcp-go"
	"sync"
)

type KcpServiceConnectList struct {
	mutexConns		sync.Mutex
	MaxConnect 		int 							//最大连接数据
	ConnectList 	map[*kcp.UDPSession]*model.PlayerConn // uid->连接Conn
}

var KcpConnPtr *KcpServiceConnectList

func InitKcpServiceConn(maxConn int){
	KcpConnPtr = &KcpServiceConnectList{
		MaxConnect:maxConn,
		ConnectList:  make(map[*kcp.UDPSession]*model.PlayerConn),
	}
}