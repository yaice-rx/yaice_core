package game

import (
	"YaIce/core"
	"YaIce/core/cluster"
	"YaIce/core/network"
	"YaIce/core/router"
	"YaIce/game/mrg"
	"YaIce/game/mrg/inside"
	"YaIce/protobuf/external"
	"YaIce/protobuf/internal_proto"
	"github.com/sirupsen/logrus"
)

type module struct {
	core.Module
}

var ModuleMrg *module = new(module)

func (this *module) RegisterRouter() {
	registerServiceRouter()
	router.RegisterRouterHandler(&c2game.C2GPing{}, mrg.PingHandler)
	router.RegisterRouterHandler(&c2game.C2GRegister{}, mrg.RegisterHandler)
}

func registerServiceRouter() { //注册内部服务
	internal_proto.RegisterServiceConnectServer(cluster.Handler.GRpcServer, &inside.Inside{})
}

func (this *module) ListenHttp() {}

func (this *module) Listen() {
	//监听端口
	if err := network.Listen(); err != nil {
		logrus.Debug(err)
		return
	}
}

func (this *module) StartHook() {

}
