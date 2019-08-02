package grpc_service

import (
	"YaIce/core/common"
	"YaIce/core/config"
	"YaIce/protobuf/internal_proto"
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"io"
	"strconv"
	"time"
)
type Client struct {
	clientConn 		*internal_proto.ServiceConnectClient
}

//启动连接GRPCService服务
func ConnectGRPCService(connectConfigData []byte)*Client{
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
	return &Client{
		clientConn:&client,
	}
}

//组装客户端发送信息
func (this *Client)SendMsg(msg interface{})error{
	var msgProtoNumber int32
	var msgData *internal_proto.MsgBodyRequest
	switch msg.(type) {
	case internal_proto.Request_ConnectStruct:
		msgData = &internal_proto.MsgBodyRequest{
			Connect: &internal_proto.Request_ConnectStruct{
			},
		}
		msgProtoNumber = common.ProtocalNumber(common.GetProtoName(&internal_proto.MsgBodyRequest{}))
		break;
	}
	data := &internal_proto.C_ServiceMsgRequest{
		MsgHandlerNumber:msgProtoNumber,
		Struct:msgData,
	}
	return registerServiceRequest(*this.clientConn,data);
}

//客户端调用
func registerServiceRequest(client internal_proto.ServiceConnectClient,r *internal_proto.C_ServiceMsgRequest)error {
	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := client.RegisterServiceRequest(ctx, r)
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