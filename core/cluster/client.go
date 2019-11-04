package cluster

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"sync"
)

type ClusterServiceModel struct {
	sync.RWMutex
	ConnMap    map[string]*grpc.ClientConn
	GRpcServer *grpc.Server
}

var Handler *ClusterServiceModel

func Init() {
	if Handler != nil {
		return
	}
	Handler = &ClusterServiceModel{
		ConnMap:    make(map[string]*grpc.ClientConn),
		GRpcServer: grpc.NewServer(),
	}
}

//启动连接GRPCService服务
func (this *ClusterServiceModel) Connect(key string, connect string) {
	conn, _ := grpc.Dial(connect,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				//当遇到此类错误，重连，否则都不予重新连接机会
				grpc_retry.WithCodes(codes.Canceled, codes.DataLoss, codes.Unavailable),
				//重连次数
				grpc_retry.WithMax(3))),
	)
	if nil != conn {
		logrus.Println("grpc 连接数据：", key, conn)
		this.ConnMap[key] = conn
	}
}

func (this *ClusterServiceModel) DeleteGRPCConn(key string) {
	logrus.Println("移除grpc连接信息：", key)
	delete(this.ConnMap, key)
}
