package inside

import (
	"YaIce/protobuf/inside_proto"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type Inside struct {
}

//处理客户端发送过来的数据
func (s *Inside) RegisterServiceRequest(r *inside_proto.C2S_Register,
	stream inside_proto.ServiceConnect_RegisterServiceRequestServer) error {
	//接收headers数据
	md, ok := metadata.FromIncomingContext(stream.Context())
	logrus.Println("接收grpc", md["key"])
	if !ok {
		logrus.Debug("metadata loading faild")
		return nil
	}
	err := stream.Send(&inside_proto.S2C_Register{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (s *Inside) SyncPlayerRequest(r *inside_proto.C2S_UserLogin,
	stream inside_proto.ServiceConnect_SyncPlayerRequestServer) error {

	return nil
}
