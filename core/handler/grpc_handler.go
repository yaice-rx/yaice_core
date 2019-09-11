package handler

import (
	"YaIce/core/config"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net"
	"strconv"
)

//连接列表
var ServerConnMap map[string]map[string]*grpc.ClientConn

//服务
var GRPCServer *grpc.Server

//初始化
func InitGPRCService() {
	GRPCServer = grpc.NewServer()
}

//初始化连接
func initClientapacity() {
	ServerConnMap = make(map[string]map[string]*grpc.ClientConn)
}

//监听grpc端口
func GRPCListen() int {
	//端口
	for port := config.CommonConfigHandler.PortStart; port <= config.CommonConfigHandler.PortEnd; port++ {
		address, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err == nil {
			go GRPCServer.Serve(address)
			return port
		}
	}
	return -1
}

//启动连接GRPCService服务
func gRPCConnInit(connect string) *grpc.ClientConn {
	conn, err := grpc.Dial(connect,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				//当遇到此类错误，重连，否则都不予重新连接机会
				grpc_retry.WithCodes(codes.Canceled, codes.DataLoss, codes.Unavailable),
				//重连次数
				grpc_retry.WithMax(3))),
	)
	if nil != err || nil == conn {
		return nil
	}
	return conn
}

func GRPCConnect(connList map[string]*grpc.ClientConn, data config.ServiceConfigModel) {
	grpcConn := gRPCConnInit(data.GetInHost() + ":" + strconv.Itoa(data.GetInPort()))
	//grpcConn := gRPCConnInit(data.GetInHost() + ":0" )
	if nil != grpcConn {
		connList[data.GetPid()] = grpcConn
	}
}

//开启连接grpc
func ConnectGRPC() {
	//connect  service
	for _, serverName := range config.ConfServiceHandler.GetConnList() {
		nodePath := config.ConfServiceHandler.GetGroupId() + "/" + serverName
		nodeData := GetEtcdNodeData(nodePath)
		if len(nodeData) <= 0 {
			return
		}
		connList := make(map[string]*grpc.ClientConn)
		for _, data := range nodeData {
			GRPCConnect(connList, data)
		}
		ServerConnMap[serverName] = connList
	}
}

func DeleteGRPCConn(name string, key string) {
	delete(ServerConnMap[name], key)
}
