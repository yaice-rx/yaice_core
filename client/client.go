package client

import (
	"YaIce/core/common"
	"YaIce/core/config"
	"YaIce/core/etcd_service"
	"YaIce/core/job"
	"YaIce/protobuf/external"
	"YaIce/protobuf/internal_proto"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"time"
)

var conn *kcp.UDPSession

func Initialize() {
	etcd_service.Init("YaIce_Service", config.GetYamlData().EtcdConnectString)

	Client := internal_proto.NewServiceConnectClient(etcd_service.EtcdClient.ConnServiceList["1/game"].Connect.ClientConn)

	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	data := &internal_proto.C_ServiceMsgRequest{
		MsgHandlerNumber: rand.Int31(),
	}

	resp, err := Client.RegisterServiceRequest(ctx, data)

	it := &internal_proto.S_ServiceMsgReply{}

	_ = resp.RecvMsg(it)

	logrus.Println("---------------------", it.MsgHandlerNumber)

	kcpconn, err := kcp.DialWithOptions("127.0.0.1:20001", nil, 10, 1)

	defer kcpconn.Close()
	if err != nil {
		fmt.Println("kcp err", err.Error())
		return
	}
	conn = kcpconn
	job.JoinJob(1, pingHandler)
	go handleKcpConn(conn)
	select {}
}

func pingHandler() {
	fmt.Println("-!-!-!")
	gmCommand := c2game.C2GGmCommand{Command: "test", Params: []string{"2312312"}}
	data, err := proto.Marshal(&gmCommand)
	if err != nil {
		log.Fatalln("Marshal client data error: ", err)
	}
	SendMsg(conn, common.ProtocalNumber("c2g_gm_command"), data)
}

func SendMsg(conn *kcp.UDPSession, protoNumber int32, data []byte) {
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
