package game

import (
	"YaIce/core"
	"YaIce/core/job"
	"YaIce/core/network"
	"YaIce/core/router"
	"YaIce/game/mrg"
	"YaIce/protobuf/outer_proto"
	"github.com/sirupsen/logrus"
)

type module struct {
	core.Module
}

var ModuleMrg *module = new(module)

func (this *module) RegisterRouter() {
	router.RegisterRouterHandler(&outer_proto.C2GPing{}, mrg.PingHandler)
	router.RegisterRouterHandler(&outer_proto.C2GLogin{}, mrg.LoginHandler)
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
