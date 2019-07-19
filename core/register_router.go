package core

import (
	"YaIce/core/common"
	"YaIce/core/connect"
	"github.com/golang/protobuf/proto"
	"sync"
)

//注册客户端请求协议
type RegisterRouterRequest struct {
	mu    sync.RWMutex
	m     map[int]func(conn *connect.PlayerConn,content []byte)
}

//注册客户端请求方法
func ( mux *RegisterRouterRequest)RegisterRouterHandler(msgObj proto.Message,handler func(conn *connect.PlayerConn,content []byte)){
	msgName := common.GetProtoName(msgObj)
	//加锁
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if mux.m == nil {
		mux.m = make(map[int]func(conn *connect.PlayerConn,content []byte))
	}
	protocolNum := common.ProtocalNumber(msgName)
	mux.m[protocolNum] = handler
}

//调用注册方法
func (mux *RegisterRouterRequest)CallRouterHandler(protoNo int,conn *connect.PlayerConn,data []byte) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	if mux.m[protoNo] != nil {
		mux.m[protoNo](conn,data)
	}
}
