package grpc_service

import (
	"YaIce/core/config"
	"encoding/json"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

//启动连接GRPCService服务
func connectGRPCService(connectConfigData []byte) {
	var model config.serviceModel
	json.Unmarshal([]byte(connectConfigData), &model)
	conn, err := grpc.Dial("127.0.0.1:20001",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				//当遇到此类错误，重连，否则都不予重新连接机会
				grpc_retry.WithCodes(codes.Canceled, codes.DataLoss, codes.Unavailable),
				//重连次数
				grpc_retry.WithMax(3))),
	)
	if nil != err {
		return nil
	}
	return &Client{
		ClientConn: conn,
	}
}

//开启连接grpc
func Connect() {
	inPort := 0
	//连接成功的时候，获取(同组)服务列表
	EtcdCli.GetNodesInfo(config.GetGroupId())
	//连接额外的服务，如果和本组相同，则不再连接
	if config.GetGroupId() == config.GetServerExtra() {
		EtcdCli.GetNodesInfo(config.GetServerExtra())
	}
	//是否是需要被其他服务器连接，不需要则不开启对外监听端口
	if config.GetIsConn() {
		//监听grpc端口
		inPort = grpc_service.ServerStart()
		if inPort == -1 {
			return -1, errors.New("grpc service start faild")
		}
	}
	return inPort, nil
}
