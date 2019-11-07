package auth

import (
	"YaIce/auth/mrg"
	"YaIce/core"
	"YaIce/core/common"
	"YaIce/core/config"
	"YaIce/core/model"
	"YaIce/core/network"
	"YaIce/core/router"
	"YaIce/protobuf/external"
	"github.com/sirupsen/logrus"
	"net/http"
)

type module struct {
	core.ModuleCore
}

var ModuleMrg *module = new(module)

func (this *module) RegisterRouter() {
	router.RegisterRouterHandler(&c2game.C2GPing{}, PingHandler)
	router.RegisterRouterHandler(&c2game.C2GRegister{}, RegisterHandler)
}

func (this *module) RegisterHook() {}

func (this *module) Listen() {
	//启动服务
	mux := http.NewServeMux()
	mux.HandleFunc("/login", mrg.Login)
	go http.ListenAndServe(":"+config.Config.HttpPort, mux)
	network.Run()
}

//处理ping包
func PingHandler(connect *model.PlayerConn, content []byte) {
	logrus.Println("ping data ", common.GetGoid())
}

func RegisterHandler(connect *model.PlayerConn, content []byte) {
	logrus.Println("register data ")
}
