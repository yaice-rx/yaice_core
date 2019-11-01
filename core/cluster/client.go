package cluster

import (
	"YaIce/core/config"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"strconv"
	"sync"
)

type ClusterServiceModel struct {
	sync.RWMutex
	ConnMap map[string]*grpc.ClientConn
	GRpcServer *grpc.Server
}

var Handler *ClusterServiceModel

func Init()*ClusterServiceModel {
	if  Handler != nil {
		return  Handler
	}
	Handler = &ClusterServiceModel{
		ConnMap:make(map[string]*grpc.ClientConn),
		GRpcServer : grpc.NewServer(),
	}
	port := Handler.gRPCListen()
	if port < 0 {
		return nil
	}
	config.StartupConfigMrg.InPort = strconv.Itoa(port)
	return Handler
}

//启动连接GRPCService服务
func (this *ClusterServiceModel)Connect(key string,connect string) {
	conn, _ := grpc.Dial(connect,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				//当遇到此类错误，重连，否则都不予重新连接机会
				grpc_retry.WithCodes(codes.Canceled, codes.DataLoss, codes.Unavailable),
				//重连次数
				grpc_retry.WithMax(3))),
	)
	if nil != conn{
		this.ConnMap[key] = conn
	}
}


func (this *ClusterServiceModel)DeleteGRPCConn( key string) {
	delete(this.ConnMap,key)
}
