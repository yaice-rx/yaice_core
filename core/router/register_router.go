package router

import (
	"YaIce/core/common"
	"YaIce/core/model"
	"github.com/golang/protobuf/proto"
	"sync"
)

//注册客户端请求协议
type RouterList struct {
	mu              sync.RWMutex
	external_member map[int32]func(conn *model.PlayerConn, content []byte)
}

var RouterListPtr *RouterList

func InitRouterList() {
	RouterListPtr = &RouterList{
		external_member: make(map[int32]func(conn *model.PlayerConn, content []byte)),
	}
}

//注册客户端请求方法
func RegisterRouterHandler(msgObj proto.Message, handler func(conn *model.PlayerConn, content []byte)) {
	msgName := common.GetProtoName(msgObj)
	//加锁
	RouterListPtr.mu.Lock()
	defer RouterListPtr.mu.Unlock()

	protocolNum := common.ProtocalNumber(msgName)
	RouterListPtr.external_member[protocolNum] = handler
}

//调用注册方法
func (mux *RouterList) CallExternalRouterHandler(protoNo int32, conn *model.PlayerConn, data []byte) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()
	if mux.external_member[protoNo] != nil {
		mux.external_member[protoNo](conn, data)
	}
}
