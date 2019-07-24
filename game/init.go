package game

import (
	"YaIce/core/config"
	"YaIce/core/etcd_service"
	"YaIce/core/grpc_service"
	"YaIce/core/kcp_service"
	"YaIce/core/router"
	"YaIce/core/temp"
	"YaIce/game/map"
	"YaIce/game/map/sort"
	"YaIce/game/mrg"
	"YaIce/protobuf/external"
	"encoding/json"
	"strconv"
)

func registerRouter(){
	router.RouterListPtr.RegisterRouterHandler(&c2game.C2GGmCommand{},mrg.CommandHandler)
	router.RouterListPtr.RegisterRouterHandler(&c2game.C2GPing{},mrg.PingHandler)
	router.RouterListPtr.RegisterRouterHandler(&c2game.C2GJoinMap{},mrg.JoinMapHandler)
}

func Initialize(){
	//-------------------------------------Init-------------------------------------//
	//加载配置文件
	temp.InitConfigData()
	//连接etcd，获取连接地址，通知网管服务器，开启地址监听
	if err := etcd_service.InitEtcd(config.ServiceConfigData.ServerGroupId,config.ServiceConfigData.ServerType);
		nil != err{
		panic("Etcd Start Failed")
		return
	}
	//注册路由
	registerRouter()
	//-------------------------------------KCP-------------------------------------//
	//监听外网端口
	ExternalPort := kcp_service.ServerExternalInit()
	if ExternalPort == -1{
		panic("All ports are occupied")
		return
	}
	//-------------------------------------End-------------------------------------//
	//-------------------------------------GRPC-------------------------------------//
	InternalPort := grpc_service.ServiceGRPCInit()
	if InternalPort == -1{
		panic("All ports are occupied")
		return
	}
	//-------------------------------------End-------------------------------------//
	//开启服务连接
	grpc_service.RegisterGateService();
	grpc_service.RegisterRelationService();
	//-------------------------------------加载路由、初始化数据-------------------------------------//
	InitServerImpl()
	//-------------------------------------ETCD 服务发现内容-------------------------------------//
	//组装自己的json
	etcdJson := etcd_service.ServerConfigEtcd{
		ServerName:		config.ServiceConfigData.ServerType,
		InternalIP:		config.ServiceConfigData.InternalHost,
		InternalPort:	strconv.Itoa(InternalPort),
		ExternalIP:		config.ServiceConfigData.ExternalHost,
		ExternalPort:	strconv.Itoa(ExternalPort),
	}
	//序列化本服务的内容
	jsonString,jsonErr := json.Marshal(etcdJson)
	if nil != jsonErr{
		panic("make json data error")
	}

	//向服务中注册自己节点数据
	etcd_service.EtcdClient.RegisterNode("",string(jsonString))
	//-------------------------------------Etcd End-------------------------------------//
	//阻塞
	select {}
}

func InitServerImpl(){
	//缓存DB数据
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
	sort.InitTown()
}




