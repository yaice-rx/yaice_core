package core

import (
	"YaIce/core/config"
	"YaIce/core/job"
	"YaIce/core/kcp_service"
	"YaIce/core/model"
	"YaIce/core/router"
	"YaIce/core/temp"
	"sync"
)

type ServerCore struct{
	CsvConfig       *temp.ConfigModule						  //程序配置文件
	MutexConns      sync.Mutex
	ServiceConfig   config.ServiceConfig                      //服务配置
	TickTasks		map[string]func()                         //tick函数列表
	DB 				*model.DBModel                           //数据库
}

var ServerCoreHandler *ServerCore

func NewServerCore()  {
	s 				:= new(ServerCore)
	//初始化路由
	router.InitRouterList();
	//初始化数据库连接
	s.DB 			= model.Init()
	//初始化网络连接信息
	kcp_service.InitKcpServiceConn(5000)
	//开启定时任务
	go job.CallJob()
	//系统核心处理
	ServerCoreHandler = s
}








