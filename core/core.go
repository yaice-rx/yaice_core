package core

import (
	"YaIce/core/config"
	"YaIce/core/dataBase"
	"YaIce/core/handler"
	"YaIce/core/job"
	"YaIce/core/kcp_service"
	"YaIce/core/router"
	"github.com/sirupsen/logrus"
	"sync"
)

type ServerCore struct {
	MutexConns sync.Mutex
	TickTasks  map[string]func() //tick函数列表
	DB         *dataBase.DBModel //数据库
}

var ServerCoreHandler *ServerCore

func NewServerCore() {
	ServerCoreHandler := new(ServerCore)
	//初始化公共配置数据
	config.InitCommonConfig()
	//初始化路由
	router.InitRouterList()
	//初始化数据库连接
	ServerCoreHandler.DB = dataBase.Connect()
	//连接Etcd
	err := handler.EtcdConnect(config.ConfServiceHandler.GetGroupId(), config.ConfServiceHandler.GetName(), config.CommonConfigHandler.EtcdConnectString)
	if nil != err {
		logrus.Debug(err.Error())
		return
	}
	//初始化grpc服务
	handler.InitGPRCService()
	//开启连接内网服务
	handler.ConnectGRPC()
	//监听内网
	port := handler.GRPCListen()
	if port == -1 {
		panic("All ports are occupied")
		return
	}
	//设置内网端口
	config.ConfServiceHandler.SetInPort(port)
	//初始化定时器
	job.Start()
	//初始化网络连接信息
	kcp_service.InitNetWork(5000)
}
