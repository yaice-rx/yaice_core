package game

import (
	"YaIce/core"
	"YaIce/core/network"
	"YaIce/core/router"
	"YaIce/game/mrg"
	"YaIce/protobuf/external"
	"github.com/sirupsen/logrus"
)

type Module struct {
	core.ModuleCore
}

var ModelMrg *Module = new(Module)

func (this *Module)RegisterRouter() {
	registerServiceRouter()
	router.RegisterRouterHandler(&c2game.C2GPing{}, mrg.PingHandler)
	router.RegisterRouterHandler(&c2game.C2GRegister{}, mrg.RegisterHandler)
}

func registerServiceRouter() {                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         //注册内部服务
	//internal_proto.RegisterServiceConnectServer(cluster.Handler.GRpcServer, &inside.Service{})
}

func (this *Module)ListenOrConnect() {
	//监听端口
	if err := network.Listen();err != nil{
		logrus.Debug(err)
		return
	}
}

//初始化数据
func (this *Module)StartHook() {
}
