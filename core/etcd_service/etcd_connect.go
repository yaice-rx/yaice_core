package etcd_service

import (
	"YaIce/core/config"
	"YaIce/core/grpc_service"
	"context"
	"encoding/json"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/sirupsen/logrus"
	"time"
)

const ttl  = 20

var EtcdClient *ClientDis

//初始化Etcd服务，
//存储结构表如下：
//serverId:1=>{"gate_序号":地址，"game_序号"：地址},
func InitEtcd(serviceName string) error{
	etcdServerList 	:=  []string{"localhost:2379"}
	etcdCli, err 	:= clientv3.New(clientv3.Config{
		Endpoints:   etcdServerList,
		DialTimeout: 5 * time.Second,
	})
	if nil != err {
		logrus.Debug("Etcd 服务，启动错误，Error Msg：",err.Error())
		return err
	}
	EtcdClient = &ClientDis{
		Endpoints:etcdServerList,
		client:etcdCli,
		serviceName:serviceName,
		path:config.ServiceConfigData.ServerGroupId+"/"+serviceName,
		ServiceList:make(map[string]*EtcdConnStruct),
	}

	if(config.ServiceConfigData.IsConnect){
		InternalPort := grpc_service.ServiceGRPCInit()
		if InternalPort == -1{
			return errors.New("grpc service start faild")
		}
		config.ServiceConfigData.InternalPort = InternalPort
	}

	//连接成功的时候，获取(同组)服务列表
	EtcdClient.GetNodesInfo(config.ServiceConfigData.ServerGroupId)
	//获取公用服务器列表
	EtcdClient.GetNodesInfo("common")
	return nil
}

//注册节点
func (c *ClientDis)RegisterNode(value string){
	if nil == c {
		return
	}
	c.grantSetLeaseKeepAlive(ttl)
	c.Lock()
	_, err := c.client.Put(context.TODO(),c.path,value,clientv3.WithLease(c.leaseRes.ID));
	if err != nil{
		logrus.Debug("数据注册失败，Error Msg：",err.Error())
		return
	}
	c.Unlock()
	go c.listenLease()
}

//授权租期，自动续约
func (c *ClientDis)grantSetLeaseKeepAlive(ttl int64) error{
	response,err :=  c.client.Lease.Grant(context.TODO(), ttl)
	if nil != err {
		return  err
	}
	c.leaseRes = response
	aliveRes,err := c.client.KeepAlive(context.TODO(),response.ID)
	if nil != err {
		return err
	}
	c.keepAliveChan = aliveRes
	return nil
}

//监测是否续约
func (c *ClientDis)listenLease(){
	for  {
		select {
		case res :=  <- c.keepAliveChan:
			if nil == res{
				logrus.Error("租期续约失败，请查看Etcd日志")
				return
			}
			break;
		}
	}
}

//获取节点数据
func (this *ClientDis)GetNodesInfo(path string){
	resp, err := this.client.Get(context.TODO(),path, clientv3.WithPrefix())
	if err != nil {
		logrus.Error("Etcd 获取内容失败")
		return
	}
	//连接服务
	for key,value := range this.extractAddrs(resp){
		//排除自己
		if key == this.path{
			continue
		}
		//不需要其他连接
		var _conf config.ServiceConfig
		json.Unmarshal([]byte(value),&_conf)
		if !_conf.IsConnect{
			continue
		}
		//连接grpc服务
		conn := grpc_service.ConnectGRPCService([]byte(value))
		if conn == nil{
			continue;
		}
		//添加服务列表中
		this.ServiceList[key] = &EtcdConnStruct{
			ConfigData:value,
			Connect:conn,
		}
	}
}

//监听节点数据变化
func (this *ClientDis) WatchNodes(key string){
	watcher := clientv3.NewWatcher(this.client)
	path := this.path+"/"+key
	for {
		rch := watcher.Watch(context.TODO(), path, clientv3.WithPrefix())
		for wresp := range rch {
			for _, event := range wresp.Events {
				switch (event.Type) {
				case mvccpb.PUT:
					//todo 维护一个列表
					var _conf config.ServiceConfig
					json.Unmarshal(event.Kv.Value,&_conf)
					if !_conf.IsConnect{
						continue
					}
					conn := grpc_service.ConnectGRPCService(event.Kv.Value)
					if nil == conn {
						continue
					}
					this.ServiceList[string(event.Kv.Key)] = &EtcdConnStruct{
						ConfigData:string(event.Kv.Value),
						Connect:conn,
					}
				case mvccpb.DELETE:
					//todo  从列表中删除
					delete(this.ServiceList, string(event.Kv.Key))
				}
			}
		}
	}
}

//读取节点数据
func (this *ClientDis) extractAddrs(resp *clientv3.GetResponse) map[string]string {
	addrs := make(map[string]string, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			key := string(resp.Kvs[i].Key)
			addrs[key] = string(v)
		}
	}
	return addrs
}

//删除节点
func (this *ClientDis)DelNode(key string)  {
	this.Lock()
	defer this.Unlock()
	path := this.path+"/"+key
	response,err := this.client.Delete(context.TODO(),path,clientv3.WithPrefix())
	if nil != err{
		logrus.Println("Error:",err.Error())
	}
	logrus.Println("Delete node",response.Deleted)
}