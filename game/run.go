package game

import (
	"YaIce/core"
	"YaIce/core/job"
	"YaIce/core/network"
	"YaIce/core/router"
	"YaIce/game/mrg"
	"YaIce/protobuf/external"
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
}

func (this *module) RegisterHook() {
	job.Crontab.AddCronTask(-1, 1, func() {
		//cluster.Send("center","auth");
	})
}

func (this *module) Listen() {
	err := network.Listen()
	//监听端口
	if err != nil {
		logrus.Debug(err)
		return
	}
	network.Run()
}
