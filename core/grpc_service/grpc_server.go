package grpc_service

import (
	"YaIce/core/config"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

var GRPCServer *grpc.Server

//监听grpc端口
func ServiceInit() int {
	//端口
	for port := config.GetYamlData().PortStart; port <= config.GetYamlData().PortEnd; port++ {
		address, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err == nil {
			go GRPCServer.Serve(address)
			return port
		}
	}
	return -1
}
