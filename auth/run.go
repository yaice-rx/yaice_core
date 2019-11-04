package auth

import (
	"YaIce/auth/mrg"
	"YaIce/auth/mrg/inside"
	"YaIce/core"
	"YaIce/core/cluster"
	"YaIce/core/config"
	"YaIce/protobuf/internal_proto"
	"net/http"
)

type module struct {
	core.ModuleCore
}

var ModuleMrg *module = new(module)

func (this *module) RegisterRouter() {
	internal_proto.RegisterServiceConnectServer(cluster.Handler.GRpcServer, &inside.Service{})
}

func (this *module) Listen() {

}

func (this *module) ListenHttp() {
	//启动服务
	mux := http.NewServeMux()
	mux.HandleFunc("/login", mrg.Login)
	http.ListenAndServe(":"+config.StartupConfigMrg.HttpPort, mux)
}

func (this *module) StartHook() {

}
