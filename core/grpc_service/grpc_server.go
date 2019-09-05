package grpc_service

import (
	"YaIce/core/config"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

var GRpcServer *grpc.Server

//初始化
func Init() {
	GRpcServer = grpc.NewServer()
}

//监听grpc端口
func Listen() int {
	//端口
	for port := config.CommonConfigData.PortStart; port <= config.CommonConfigData.PortEnd; port++ {
		address, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err == nil {
			go GRpcServer.Serve(address)
			return port
		}
	}
	return -1
}
