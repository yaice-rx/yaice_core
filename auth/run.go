package auth

import (
	"YaIce/auth/mrg"
	"YaIce/auth/mrg/inside"
	"YaIce/core"
	"YaIce/core/cluster"
	"YaIce/core/config"
	"YaIce/protobuf/inside_proto"
	"net/http"
)

type module struct {
	core.ModuleCore
}

var ModuleMrg *module = new(module)

func (this *module) RegisterRouter() {
	inside_proto.RegisterServiceConnectServer(cluster.Handler.GRpcServer, &inside.Service{})
}

func (this *module) RegisterHook() {}

func (this *module) Listen() {
	//启动服务
	mux := http.NewServeMux()
	mux.HandleFunc("/login", mrg.Login)
	http.ListenAndServe(":"+config.Config.HttpPort, mux)
}
