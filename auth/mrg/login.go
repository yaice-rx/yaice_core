package mrg

import (
	"YaIce/core/model"
	"encoding/json"
	"github.com/satori/go.uuid"
	"net/http"
)

func Login(w http.ResponseWriter, resp *http.Request) {
	//获取已经连接的设备
	session := uuid.Must(uuid.NewV4()).String()
	token := model.LoginToken{
		//Pid:        config.ConfServiceHandler.GetPid(),
		SessionKey: session,
		Host:       "10.0.0.10",
		Port:       20001,
	}
	data, _ := json.Marshal(token)
	w.Write(data)
}
