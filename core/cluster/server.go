package cluster

import (
	"YaIce/core/config"
	"YaIce/core/yaml"
	"net"
	"strconv"
)

//监听grpc端口
func (this *ClusterServiceModel) gRPCListen() int {
	//端口
	for port := yaml.YamlDevMrg.NetworkPortStart; port <= yaml.YamlDevMrg.NetworkPortEnd; port++ {
		address, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err == nil {
			go this.GRpcServer.Serve(address)
			return port
		}
	}
	return -1
}

func Listen() {
	port := Handler.gRPCListen()
	if port < 0 {
		return
	}
	config.Config.InPort = strconv.Itoa(port)
}
