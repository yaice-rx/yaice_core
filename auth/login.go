package auth

import (
	Auth_Model "YaIce/auth/model"
	"YaIce/core/model"
	"encoding/json"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

var accountLoginSessionMap map[string]Auth_Model.AccountLoginSession = make(map[string]Auth_Model.AccountLoginSession)

func Initialize(port string, server_id string) {
	//注册内部路由
	registerRouter()
	//向服务中注册自己节点数据
	//handler.RegisterServiceConfigData()
	//监听Http服务器
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	http.ListenAndServe(":"+port, mux)
	select {

	}
}

func login(w http.ResponseWriter, resp *http.Request) {
	//获取已经连接的设备
	session := uuid.Must(uuid.NewV4()).String()
	guid := time.Now().Unix()
	token := model.LoginToken{
		//Pid:        config.ConfServiceHandler.GetPid(),
		SessionKey: session,
		Host:       "10.0.0.10",
		Port:       20001,
	}
	data, _ := json.Marshal(token)
	accountLoginSessionMap[session] = Auth_Model.AccountLogin(session, guid)
	w.Write(data)
}
