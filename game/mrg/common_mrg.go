package mrg

import (
	"YaIce/core/connect"
	"YaIce/protobuf/external"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

//Gm命令处理
func  CommandHandler(conn *connect.PlayerConn,content []byte){
	fmt.Println("-----------------")
	data := c2game.C2GGmCommand{}
	c2g_proto := c2game.C2GSyncTime{ClientTime:"test current time :"+time.Now().String()}
	err := proto.Unmarshal(content,&data)
	if err != nil {
		log.Fatalln("Unmarshal data error: ", err)
	}
	conn.WriteMsg(&c2g_proto)
}

//处理ping包
func  PingHandler(connect *connect.PlayerConn,content []byte)  {
	data := c2game.C2GPing{}
	err := proto.Unmarshal(content,&data)
	if err != nil {
		log.Fatalln("Unmarshal data error: ", err)
	}
}