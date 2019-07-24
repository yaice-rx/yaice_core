package client

import (
	"YaIce/core/common"
	"YaIce/core/etcd_service"
	"YaIce/core/job"
	"YaIce/protobuf/external"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	"log"
	"time"
)

var conn *kcp.UDPSession

func Initialize(){
	etcd_service.InitEtcd("1","game")
	//etcdCli.RegisterNode("1","test-=-=-=-=")
	time.Sleep(2 * time.Second)
	data,_ := etcd_service.EtcdClient.GetNodesInfo("")
	fmt.Println(data)
	etcd_service.EtcdClient.WatchNodes("");


	kcpconn, err := kcp.DialWithOptions("127.0.0.1:20001", nil, 10, 1)
	defer  kcpconn.Close()
	if err != nil {
		fmt.Println("kcp err",err.Error())
		return
	}
	conn = kcpconn
	job.JoinJob(1,pingHandler)
	go handleKcpConn(conn)
	select {}
}

func  pingHandler(){
	fmt.Println("-!-!-!")
	gmCommand := c2game.C2GGmCommand{Command:"test",Params:[]string{"2312312"}}
	data, err := proto.Marshal(&gmCommand)
	if err != nil {
		log.Fatalln("Marshal client data error: ", err)
	}
	SendMsg(conn, common.ProtocalNumber("c2g_gm_command"),data)
}


func SendMsg(conn *kcp.UDPSession,protoNumber int,data []byte) {
	if conn != nil {
		content := common.IntToBytes(protoNumber)
		content = append(content, data...)
		_, err := conn.Write(content)
		if err != nil {
			fmt.Println("send msg error ", err.Error())
		}
	} else {
		fmt.Println("kcp connect is nil ", time.Now().String())
	}
}


func handleKcpConn(conn *kcp.UDPSession) {
 	buf := make([]byte, 65535)
	for {
		num, err := conn.Read(buf)
		if err != nil {
			fmt.Println("接收数据失败!", err)
			return
		}
		fmt.Printf("接收服务端数据长度：%d, 数据：%s\n", num, string(buf[4:num]))
		time.Sleep(time.Second)
	}
}