package mrg

import (
	"YaIce/core/handler"
	"YaIce/core/model"
	"YaIce/protobuf/external"
	"YaIce/protobuf/internal_proto"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"time"
)

//处理ping包
func PingHandler(connect *model.PlayerConn, content []byte) {
	logrus.Println("=========ping=============", time.Now().String())
}

func RegisterHandler(connect *model.PlayerConn, content []byte) {
	data := c2game.C2GRegister{}
	err := proto.Unmarshal(content, &data)
	if err != nil {
		logrus.Println("Unmarshal data error: ", err)
	}
	logrus.Println(handler.ServerMapHandler["auth"][data.Pid])
	//向auth服务器发起请求
	Client := internal_proto.NewServiceConnectClient(handler.ServerMapHandler["auth"][data.Pid])

	_, err = Client.RegisterServiceRequest(context.Background(), &internal_proto.C2S_Register{})
	if nil != err {
		logrus.Println("logrus error :", err.Error())
	}

	logrus.Println("==========register============", data.SessionId, data.Pid)
}
