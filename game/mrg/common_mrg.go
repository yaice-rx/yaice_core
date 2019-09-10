package mrg

import (
	"YaIce/core/model"
	"YaIce/protobuf/external"
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
	logrus.Println("==========register============", data.SessionId, data.Pid)
}
