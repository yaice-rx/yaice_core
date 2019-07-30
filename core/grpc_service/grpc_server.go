package grpc_service

import (
	"YaIce/core/temp"
	"YaIce/protobuf/internal_proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strconv"
)


type serverConns struct {
	conns map[string]*internal_proto.ServiceConnect_RegisterServiceRequestServer
}

var serverConnList *serverConns

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
	stream internal_proto.ServiceConnect_RegisterServiceRequestServer)error  {
	if nil == serverConnList.conns[r.Header.Uid] {
		serverConnList.conns[r.Header.Uid] = &stream
	}
	//todo  处理client连接
	err := stream.Send(&internal_proto.S_ServiceMsgReply{MsgHandlerNumber:123})
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
