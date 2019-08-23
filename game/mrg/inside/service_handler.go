package inside

import (
	"YaIce/protobuf/internal_proto"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type Service struct {
}

//处理客户端发送过来的数据
func (s *Service) RegisterServiceRequest(r *internal_proto.C_ServiceMsgRequest,
	stream internal_proto.ServiceConnect_RegisterServiceRequestServer) error {
	//接收headers数据
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		logrus.Debug("metadata loading faild")
		return nil
	}
	_ = md.Get(":authority")[0]

	logrus.Println("revice value grpc ...", r.MsgHandlerNumber)
	err := stream.Send(&internal_proto.S_ServiceMsgReply{MsgHandlerNumber: r.MsgHandlerNumber})
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (s *Service) SyncPlayerRequest(r *internal_proto.C2S_Register,
	stream internal_proto.ServiceConnect_SyncPlayerRequestServer) error {
	return nil
}
