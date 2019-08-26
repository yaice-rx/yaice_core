package grpc_service

import (
	"YaIce/core/config"
	"encoding/json"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Client struct {
	ClientConn *grpc.ClientConn
}

//启动连接GRPCService服务
func ConnectGRPCService(connectConfigData []byte) *Client {
	var model config.ServiceModel
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
