package grpc_service

import (
	"YaIce/core/config"
	"YaIce/protobuf/internal_proto"
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"io"
	"strconv"
)

//启动连接GRPCService服务
func ConnectGRPCService(connectConfigData []byte)*internal_proto.ServiceConnectClient{
	var serviceConfig config.ServiceConfig
	json.Unmarshal([]byte(connectConfigData),&serviceConfig)
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
		return nil
	}
	client := internal_proto.NewServiceConnectClient(conn)
	return &client
}

//客户端调用
func RegisterServiceRequest(client internal_proto.ServiceConnectClient,r *internal_proto.C_ServiceMsgRequest)error {
	stream, err := client.RegisterServiceRequest(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		//todo  处理服务器的消息
		logrus.Println("grpc client receive",resp.MsgHandlerNumber)
	}
	return nil
}