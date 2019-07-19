package game

import (
	"YaIce/core"
	"YaIce/core/connect"
	"YaIce/core/temp"
	"YaIce/game/map"
	"YaIce/game/map/sort"
	"YaIce/game/mrg"
	"YaIce/protobuf"
	"github.com/sirupsen/logrus"
)

func register(router *core.RegisterRouterRequest){
	router.RegisterRouterHandler(&c2game.C2GGmCommand{},mrg.CommandHandler)
	router.RegisterRouterHandler(&c2game.C2GPing{},mrg.PingHandler)
	router.RegisterRouterHandler(&c2game.C2GJoinMap{},mrg.JoinMapHandler)
}
func Initialize(core *core.ServerCore,server_id string){
	//连接etcd，获取连接地址，通知网管服务器，开启地址监听
	etcdCli,_ := connect.InitEtcd(server_id,core.ServerType)
	//加载配置文件
	temp.InitConfigData()
	//注册路由
	register(core.Routers)
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
	//监听外网端口
	core.ServerExternalInit()
	//组装自己的json
	etcdJson := connect.ServerConfigEtcd{
		ServerName:core.ServerType,
		InternalIP:core.InternalHost,
		ExternalIP:core.ExternalHost,
	}
	content,err := etcdCli.GetNodesInfo("")
	if err != nil{
		logrus.Debug(err.Error())
		return
	}
	logrus.Println(content)
	go etcdCli.WatchNodes("0")
}




