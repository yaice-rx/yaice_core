package mrg

import (
	"YaIce/core/model"
	"YaIce/protobuf/external"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

//Gm命令处理
func CommandHandler(conn *model.PlayerConn, content []byte) {
	log.Println("g2c gm command data")
	data := c2game.C2GGmCommand{}
	c2g_proto := c2game.C2GSyncTime{ClientTime: "test current time :" + time.Now().String()}
	err := proto.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln("Unmarshal data error: ", err)
	}
	conn.WriteMsg(&c2g_proto)
}

//处理ping包
func PingHandler(connect *model.PlayerConn, content []byte) {
	data := c2game.C2GPing{}
	err := proto.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln("Unmarshal data error: ", err)
	}
}

func RegisterHandler(connect *model.PlayerConn, content []byte) {

}
