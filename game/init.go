package game

import (
	"YaIce/conf"
	"YaIce/core/config"
	"YaIce/core/handler"
	"YaIce/core/kcp_service"
	"YaIce/core/router"
	"YaIce/game/mrg"
	"YaIce/game/mrg/inside"
	"YaIce/protobuf/external"
	"YaIce/protobuf/internal_proto"
	"encoding/json"
)

func registerRouter() {
	registerServiceRouter()
	router.RegisterRouterHandler(&c2game.C2GGmCommand{}, mrg.CommandHandler)
	router.RegisterRouterHandler(&c2game.C2GPing{}, mrg.PingHandler)
	router.RegisterRouterHandler(&c2game.C2GRegister{}, mrg.RegisterHandler)
	router.RegisterRouterHandler(&c2game.C2GJoinMap{}, mrg.JoinMapHandler)
}

func registerServiceRouter() {
	internal_proto.RegisterServiceConnectServer(handler.GRPCServer, &inside.Service{})
}

func Initialize() {
	//注册路由
	registerRouter()
	//监听外网端口
	port := kcp_service.Listen()
	if port == -1 {
		panic("All ports are occupied")
		return
	}
	//设置外网监听端口
	config.ConfServiceHandler.SetOutPort(port)
	//监听内网
	port = handler.GRPCListen()
	if port == -1 {
		panic("All ports are occupied")
		return
	}
	//设置内网端口
	config.ConfServiceHandler.SetInPort(port)
	//序列化本服务的内容
	jsonString, jsonErr := json.Marshal(config.ConfServiceHandler.GetServiceConfData())
	if nil != jsonErr {
		panic("make json data error")
	}
	//向服务中注册自己节点数据
	handler.RegisterEtcdData(string(jsonString))
	//-------------------------------------End-------------------------------------//
	//初始化配置
	InitServerImpl()
	//启动服务
	kcp_service.Run()
}

//初始化数据
func InitServerImpl() {
	//开启连接内网服务
	handler.ConnectGRPC()
	//初始化CSV配置文件数据
	conf.InitCSVConfigData()
	/*//缓存DB数据
	mrg.InitCacheDBData()
	//初始化地形
	_map.InitTerrain()
	//初始化视野
	_map.InitVision()
	//初始化野怪
	sort.InitMonster()
	//初始化资源
	sort.InitResource()
	//初始化城市
	sort.InitTown()*/
}
