package etcd_service

import (
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"sync"
)

type ServerConfigEtcd struct {
	ServerName 	 	string `json:"server_name"`
	InternalIP   	string `json:"internal_ip"`
	InternalPort   	string `json:"internal_port"`
	ExternalIP   	string `json:"external_ip"`
	ExternalPort  	string `json:"external_port"`
}

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
	ServiceList  	map[string]string
	leaseRes    	*clientv3.LeaseGrantResponse
	keepAliveChan  	<-chan *clientv3.LeaseKeepAliveResponse
}