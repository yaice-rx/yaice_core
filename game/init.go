package game

import (
	"YaIce/core"
	"YaIce/core/connect"
	"YaIce/core/temp"
	"YaIce/game/map"
	"YaIce/game/map/sort"
	"YaIce/game/mrg"
	"YaIce/protobuf"
	"encoding/json"
	"fmt"
	"strconv"
)

func register(router *core.RegisterRouterRequest){
	router.RegisterRouterHandler(&c2game.C2GGmCommand{},mrg.CommandHandler)
	router.RegisterRouterHandler(&c2game.C2GPing{},mrg.PingHandler)
	router.RegisterRouterHandler(&c2game.C2GJoinMap{},mrg.JoinMapHandler)
}
func Initialize(core *core.ServerCore,server_id string){
	//加载配置文件
	temp.InitConfigData()
	//注册路由
	register(core.Routers)
	//连接etcd，获取连接地址，通知网管服务器，开启地址监听
	etcdCli,_ := connect.InitEtcd(server_id,core.ServerType)

	//监听外网端口
	ExternalPort := core.ServerExternalInit()
	fmt.Println("监听外网端口：",ExternalPort)
	if ExternalPort == -1{
		panic("All ports are occupied")
		return
	}

	InternalPort := core.ServerInternalInit()
	fmt.Println("监听内网端口：",InternalPort)
	if InternalPort == -1{
		panic("All ports are occupied")
		return
	}
	//组装自己的json
	etcdJson := connect.ServerConfigEtcd{
		ServerName:core.ServerType,
		InternalIP:core.InternalHost,
		InternalPort:strconv.Itoa(InternalPort),
		ExternalIP:core.ExternalHost,
		ExternalPort:strconv.Itoa(ExternalPort),
	}
	//序列化本服务的内容
	jsonString,jsonErr := json.Marshal(etcdJson)
	if nil != jsonErr{
		panic("make json data error")
	}
	fmt.Println(string(jsonString))
	//向etcd注册服务内容
	etcdCli.RegisterNode("",string(jsonString))
	//-------------------------------------加载路由、初始化数据-------------------------------------------------//

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




