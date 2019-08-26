package game

import (
	"YaIce/conf"
	"YaIce/core/config"
	"YaIce/core/etcd_service"
	"YaIce/core/grpc_service"
	"YaIce/core/kcp_service"
	"YaIce/core/router"
	"YaIce/game/mrg"
	"YaIce/game/mrg/inside"
	"YaIce/protobuf/external"
	"YaIce/protobuf/internal_proto"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func registerRouter() {
	registerServiceRouter()
	router.RegisterRouterHandler(&c2game.C2GGmCommand{}, mrg.CommandHandler)
	router.RegisterRouterHandler(&c2game.C2GPing{}, mrg.PingHandler)
	router.RegisterRouterHandler(&c2game.C2GRegister{}, mrg.RegisterHandler)
	router.RegisterRouterHandler(&c2game.C2GJoinMap{}, mrg.JoinMapHandler)
}

func registerServiceRouter() {
	internal_proto.RegisterServiceConnectServer(grpc_service.GRPCServer, &inside.Service{})
}

func Initialize() {
	//注册路由
	registerRouter()
	//监听外网端口
	port := kcp_service.Init()
	if port == -1 {
		panic("All ports are occupied")
		return
	}
	//连接Etcd服务,连接服务
	inPort, err := etcd_service.Init(config.GetName(), config.GetYamlData().EtcdConnectString)
	if nil != err || inPort == -1 {
		logrus.Debug(err.Error())
		return
	}
	//序列化本服务的内容
	jsonString, jsonErr := json.Marshal(config.GetServiceConfData())
	if nil != jsonErr {
		panic("make json data error")
	}
	//向服务中注册自己节点数据
	etcd_service.EtcdClient.RegisterNode(string(jsonString))
	//-------------------------------------End-------------------------------------//
	//初始化配置
	InitServerImpl()
	//启动服务
	kcp_service.Run()
}

//初始化数据
func InitServerImpl() {
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
