package etcd_service

import (
	"YaIce/protobuf/internal_proto"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"sync"
)

type EtcdConnStruct struct {
	ConfigData  string
	ConnectName string
	Connect		*internal_proto.ServiceConnectClient
}

type ClientDis struct {
	sync.RWMutex
	client 			*clientv3.Client
	Endpoints   	[]string						//连接Etcd服务列表
	serviceName		string							//监听服务名称
	path 			string									//服务在Etcd中的key
	Header			*internal_proto.Request_HeaderStruct	//grpc请求协议头
	ConnServiceList  	map[string]*EtcdConnStruct		//连接的grpc列表
	leaseRes    	*clientv3.LeaseGrantResponse	//自己配置租约
	keepAliveChan  	<-chan *clientv3.LeaseKeepAliveResponse
	LocalServer 	*grpc.Server					//自己grpc服务
}