package grpc_service

import (
	"YaIce/core/config"
	"YaIce/core/etcd_service"
	"YaIce/core/temp"
	"encoding/json"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

//监听grpc端口
func ServiceGRPCInit()int{
	//从zookeeper中获取登陆服务器的ip
	server := grpc.NewServer()
	etcd_service.EtcdClient.LocalServer = server
	//注册路由
	reflection.Register(server)
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

//启动连接GRPCService服务
func ConnectGRPCService(connectConfigData []byte)*grpc.ClientConn{
	var serviceConfig config.ServiceConfig
	json.Unmarshal(connectConfigData,&serviceConfig)
	conn, err := grpc.Dial(serviceConfig.InternalHost+":"+strconv.Itoa(serviceConfig.InternalPort),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				//当遇到此类错误，重连，否则都不予重新连接机会
				grpc_retry.WithCodes(codes.Canceled,codes.DataLoss,codes.Unavailable),
				//重连次数
				grpc_retry.WithMax(3))),
	)
	if nil != err{
		logrus.Error("服务连接IP："+serviceConfig.InternalHost+"失败，Error Msg :",err.Error())
		return nil
	}
	return conn
}
