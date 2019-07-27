package etcd_service

import (
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"sync"
)

type EtcdConnStruct struct {
	ConfigData  string
	Connect		*grpc.ClientConn
}

type ClientDis struct {
	sync.RWMutex
	client 			*clientv3.Client
	Endpoints   	[]string						//连接Etcd服务列表
	serviceName		string							//服务名称
	path 			string							//自己配置路径
	ServiceList  	map[string]*EtcdConnStruct		//连接的grpc列表
	leaseRes    	*clientv3.LeaseGrantResponse	//自己配置租约
	keepAliveChan  	<-chan *clientv3.LeaseKeepAliveResponse
	LocalServer 	*grpc.Server					//自己grpc服务
}