package inside

import (
	"YaIce/protobuf/inside_proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"sync"
)

type Service struct {
	mu sync.Mutex
}

//处理客户端发送过来的数据
func (s *Service) RegisterServiceRequest(r *inside_proto.C2S_Register,
	stream inside_proto.ServiceConnect_RegisterServiceRequestServer) error {
	err := stream.Send(&inside_proto.S2C_Register{})
	md, _ := metadata.FromIncomingContext(stream.Context())
	logrus.Println("接收grpc", md["key"])
	if err != nil {
		logrus.Println(err.Error())
	}
	return nil
}

func (s *Service) SyncPlayerRequest(r *inside_proto.C2S_UserLogin,
	stream inside_proto.ServiceConnect_SyncPlayerRequestServer) error {
	//接收headers数据
	_, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		logrus.Debug("metadata loading faild")
		return nil
	}
	return nil
}
