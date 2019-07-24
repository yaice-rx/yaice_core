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
	m     map[int]func(conn *model.PlayerConn,content []byte)
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

	if mux.m == nil {
		mux.m = make(map[int]func(conn *model.PlayerConn,content []byte))
	}
	protocolNum := common.ProtocalNumber(msgName)
	mux.m[protocolNum] = handler
}

//调用注册方法
func (mux *RouterList)CallRouterHandler(protoNo int,conn *model.PlayerConn,data []byte) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()
	if mux.m[protoNo] != nil {
		mux.m[protoNo](conn,data)
	}
}
