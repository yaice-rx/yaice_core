package auth

import (
	"YaIce/core/config"
	"YaIce/core/etcd_service"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Initialize(port string, server_id string) {
	//注册内部路由
	registerRouter()
	//连接Etcd服务,连接服务
	inPort, err := etcd_service.Init(config.GetName(), config.GetYamlData().EtcdConnectString)
	if nil != err || inPort == -1 {
		logrus.Debug(err.Error())
		return
	}
	config.SetInPort(inPort)
	//序列化本服务的内容
	jsonString, jsonErr := json.Marshal(config.GetServiceConfData())
	if nil != jsonErr {
		panic("make json data error")
	}
	//向服务中注册自己节点数据
	etcd_service.RegisterNode(string(jsonString))
	//监听Http服务器
	mux := http.NewServeMux()
	mux.HandleFunc("/", login)
	http.ListenAndServe(":"+port, mux)
}

func login(w http.ResponseWriter, r *http.Request) {
	logrus.Println(r.RequestURI)
}
