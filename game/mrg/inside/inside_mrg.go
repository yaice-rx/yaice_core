package inside

import (
	"YaIce/protobuf/internal_proto"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type Inside struct {
}

//处理客户端发送过来的数据
func (s *Inside) RegisterServiceRequest(r *internal_proto.C2S_Register,
	stream internal_proto.ServiceConnect_RegisterServiceRequestServer) error {
	//接收headers数据
	_, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		logrus.Debug("metadata loading faild")
		return nil
	}
	err := stream.Send(&internal_proto.S2C_Register{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (s *Inside) SyncPlayerRequest(r *internal_proto.C2S_UserLogin,
	stream internal_proto.ServiceConnect_SyncPlayerRequestServer) error {

	return nil
}
