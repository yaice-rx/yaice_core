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
	Endpoints   	[]string
	serviceName		string
	path 			string
	ServiceList  	map[string]*EtcdConnStruct
	leaseRes    	*clientv3.LeaseGrantResponse
	keepAliveChan  	<-chan *clientv3.LeaseKeepAliveResponse
}