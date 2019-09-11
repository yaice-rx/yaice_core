package router

import (
	"YaIce/core/common"
	"YaIce/core/model"
	"github.com/golang/protobuf/proto"
	"sync"
)

//注册客户端请求协议
type RouterList struct {
	mu           sync.RWMutex
	outerFuncMap map[int32]func(conn *model.PlayerConn, content []byte)
	filterMap    []func(conn *model.PlayerConn, content []byte)
}

var routerListPtr *RouterList

func InitRouterList() {
	routerListPtr = &RouterList{
		outerFuncMap: make(map[int32]func(conn *model.PlayerConn, content []byte)),
	}
}

//注册客户端请求方法
func RegisterRouterHandler(msgObj proto.Message, handler func(conn *model.PlayerConn, content []byte)) {
	//加锁
	routerListPtr.mu.Lock()
	defer routerListPtr.mu.Unlock()
	msgName := common.GetProtoName(msgObj)
	protocolNum := common.ProtocalNumber(msgName)
	routerListPtr.outerFuncMap[protocolNum] = handler
}

//调用注册方法
func CallExternalRouterHandler(protoNo int32, conn *model.PlayerConn, data []byte) {
	routerListPtr.mu.RLock()
	defer routerListPtr.mu.RUnlock()
	if routerListPtr.outerFuncMap[protoNo] != nil {
		routerListPtr.outerFuncMap[protoNo](conn, data)
	}
}

//注册过滤处理
func RegisterFilterHandler(handler func(conn *model.PlayerConn, content []byte)) {
	//加锁
	routerListPtr.mu.Lock()
	defer routerListPtr.mu.Unlock()
	routerListPtr.filterMap = append(routerListPtr.filterMap, handler)
}

//注册过滤处理
func CallFilterHandler(conn *model.PlayerConn, data []byte) {
	//加锁
	routerListPtr.mu.Lock()
	defer routerListPtr.mu.Unlock()
	for _, handler := range routerListPtr.filterMap {
		handler(conn, data)
	}
}

func RegisterInterServiceHandler(handler interface{}) {

}
