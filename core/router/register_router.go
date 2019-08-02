package router

import (
	"YaIce/core/common"
	"YaIce/core/model"
	"github.com/golang/protobuf/proto"
	"sync"
)

//注册客户端请求协议
type RouterList struct {
	mu    sync.RWMutex
	external_member     map[int32]func(conn *model.PlayerConn,content []byte)
	internal_member     map[int32]func(content[]byte)
}

var RouterListPtr   *RouterList

func InitRouterList(){
	RouterListPtr = &RouterList{}
}

//注册客户端请求方法
func ( mux *RouterList)RegisterRouterHandler(msgObj proto.Message,handler func(conn *model.PlayerConn,content []byte)){
	msgName := common.GetProtoName(msgObj)
	//加锁
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if mux.external_member == nil {
		mux.external_member = make(map[int32]func(conn *model.PlayerConn,content []byte))
	}
	protocolNum := common.ProtocalNumber(msgName)
	mux.external_member[protocolNum] = handler
}

func ( mux *RouterList)RegisterInternalRouterHandler(msgObj proto.Message,handler func(content[]byte)){
	msgName := common.GetProtoName(msgObj)
	//加锁
	mux.mu.Lock()
	defer mux.mu.Unlock()

	protocolNum := common.ProtocalNumber(msgName)
	mux.internal_member[protocolNum] = handler
}

//调用注册方法
func (mux *RouterList)CallExternalRouterHandler(protoNo int32,conn *model.PlayerConn,data []byte) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()
	if mux.external_member[protoNo] != nil {
		mux.external_member[protoNo](conn,data)
	}
}

//调用注册方法
func (mux *RouterList)CallInternalRouterHandler(protoNo int32,conn *model.PlayerConn,data []byte) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()
	if mux.internal_member[protoNo] != nil {
		mux.internal_member[protoNo](data)
	}
}
