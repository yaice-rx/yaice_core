package core

import (
	"YaIce/core/grpc_service"
	"YaIce/core/job"
	"YaIce/core/kcp_service"
	"YaIce/core/model"
	"YaIce/core/router"
	"sync"
)

type ServerCore struct {
	MutexConns sync.Mutex
	TickTasks  map[string]func() //tick函数列表
	DB         *model.DBModel    //数据库
}

var ServerCoreHandler *ServerCore

func NewServerCore() {
	s := new(ServerCore)
	//初始化路由
	router.InitRouterList()
	//初始化数据库连接
	s.DB = model.Init()
	//初始化网络连接信息
	kcp_service.InitKcpServiceConn(5000)
	//初始化grpc服务
	grpc_service.Init()
	//开启定时任务
	go job.CallJob()
	//系统核心处理
	ServerCoreHandler = s
}
