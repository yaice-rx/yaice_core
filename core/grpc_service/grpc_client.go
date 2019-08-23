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
	"google.golang.org/grpc/metadata"
	"io"
	"time"
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

//组装客户端发送信息
func (this *Client) SendMsg(msg interface{}) error {
	/*var msgProtoNumber int32
	var msgData *internal_proto.C2S_Body
	switch msg.(type) {
	case internal_proto.C2S_Register:
		msgData = &internal_proto.C2S_Body{
			Register: &internal_proto.C2S_Register{
			},
		}
		msgProtoNumber = common.ProtocalNumber(common.GetProtoName(&internal_proto.C2S_Register{}))
		break;
	}
	data := &internal_proto.C_ServiceMsgRequest{
		MsgHandlerNumber:msgProtoNumber,
		Body:msgData,
	}*/
	//return registerServiceRequest(*this.ClientConn,data);
	return nil
}

//客户端调用
func registerServiceRequest(client internal_proto.ServiceConnectClient, r *internal_proto.C_ServiceMsgRequest) error {
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
		logrus.Println("grpc client receive", resp.MsgHandlerNumber)
	}
	return nil
}
