package grpc_service

import (
	"YaIce/core/temp"
	"YaIce/protobuf/internal_proto"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"strconv"
)

type serverConns struct {
	conns map[string]*internal_proto.ServiceConnect_RegisterServiceRequestServer
}

var serverConnList *serverConns

var grpc_conn_list_key string ;

//监听grpc端口
func ServiceGRPCInit()int{
	//从zookeeper中获取登陆服务器的ip
	server := grpc.NewServer()
	//初始化连接
	serverConnList = &serverConns{
		conns :make(map[string]*internal_proto.ServiceConnect_RegisterServiceRequestServer),
	}
	internal_proto.RegisterServiceConnectServer(server, &ServiceRegister{})
	//获取 端口
	for port := temp.ConfigCacheData.YamlConfigData.PortStart; port <= temp.ConfigCacheData.YamlConfigData.PortEnd; port++{
		address, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err == nil {
			go server.Serve(address)
			return port;
		}
	}
	return -1
}

type ServiceRegister struct {

}

//处理客户端发送过来的数据
func (s *ServiceRegister)RegisterServiceRequest( r *internal_proto.C_ServiceMsgRequest,
	stream internal_proto.ServiceConnect_RegisterServiceRequestServer) error {
	//接收headers数据
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		logrus.Debug("metadata loading faild")
		return nil
	}
	grpc_conn_list_key = md.Get(":authority")[0]

	if nil == serverConnList.conns[grpc_conn_list_key] {
		serverConnList.conns[grpc_conn_list_key] = &stream
	}

	//todo  处理client连接
	//router.RouterListPtr.CallInternalRouterHandler()

	err := stream.Send(&internal_proto.S_ServiceMsgReply{MsgHandlerNumber:123})
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
