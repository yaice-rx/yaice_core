package game

import (
	"YaIce/core"
	"YaIce/core/cluster"
	"YaIce/core/job"
	"YaIce/core/network"
	"YaIce/core/router"
	"YaIce/game/mrg"
	"YaIce/game/mrg/inside"
	"YaIce/protobuf/external"
	"YaIce/protobuf/inside_proto"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
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
	inside_proto.RegisterServiceConnectServer(cluster.Handler.GRpcServer, &inside.Inside{})
}

func (this *module) RegisterHook() {
	job.Crontab.AddCronTask(10, -1, func() {
		for _, value := range cluster.Handler.ConnMap {
			md := metadata.AppendToOutgoingContext(context.TODO(), "key", "234234")
			conn := inside_proto.NewServiceConnectClient(value)
			_, err := conn.RegisterServiceRequest(md, &inside_proto.C2S_Register{})
			if err != nil {
				logrus.Println("could not greet: %v", err)
			}
			logrus.Println("Greeting: %s")
		}
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
