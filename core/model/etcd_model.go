package model

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"sync"
)

type ClientModel struct {
	sync.RWMutex
	context.Context
	Client        *clientv3.Client
	Endpoints     []string                     //连接Etcd服务列表
	ServiceName   string                       //监听服务名称
	Path          string                       //服务在Etcd中的key
	LeaseRes      *clientv3.LeaseGrantResponse //自己配置租约
	KeepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	LocalServer   *grpc.Server //自己grpc服务
}
