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

//初始化Etcd服务，
//存储结构表如下：
//serverId:1=>{"gate_序号":地址，"game_序号"：地址},
func Init(serviceName string) (int, error) {
	inPort := -1
	serverList := []string{"localhost:2379"}
	//连接etcd服务
	client, err := clientv3.New(clientv3.Config{Endpoints: serverList, DialTimeout: 5 * time.Second})
	if nil != err {
		logrus.Debug("Etcd 服务，启动错误，Error Msg：", err.Error())
		return -1, err
	}
	//初始化连接信息
	EtcdClient = &ClientDis{
		Endpoints:       serverList,
		client:          client,
		serviceName:     serviceName,
		path:            config.GetGroupId() + "/" + serviceName,
		ConnServiceList: make(map[string]*EtcdConnStruct),
	}

	//连接成功的时候，获取(同组)服务列表
	EtcdClient.GetNodesInfo(config.GetGroupId())
	//连接额外的服务，如果和本组相同，则不再连接
	if config.GetGroupId() == config.GetServerExtra() {
		EtcdClient.GetNodesInfo(config.GetServerExtra())
	}
	//是否是需要被其他服务器连接，不需要则不开启对外监听端口
	if config.GetIsConn() {
		//监听grpc端口
		inPort = grpc_service.ServiceInit()
		if inPort == -1 {
			return -1, errors.New("grpc service start faild")
		}
	}
	config.SetInPort(inPort)
	return inPort, nil
}

//注册节点
func (c *ClientDis) RegisterNode(value string) {
	if nil == c {
		return
	}
	c.grantSetLeaseKeepAlive(ttl)
	c.Lock()
	_, err := c.client.Put(context.TODO(), c.path, value, clientv3.WithLease(c.leaseRes.ID))
	if err != nil {
		logrus.Debug("数据注册失败，Error Msg：", err.Error())
		return
	}
	c.Unlock()
	go c.listenLease()
}

//授权租期，自动续约
func (c *ClientDis) grantSetLeaseKeepAlive(ttl int64) error {
	response, err := c.client.Lease.Grant(context.TODO(), ttl)
	if nil != err {
		return err
	}
	c.leaseRes = response
	aliveRes, err := c.client.KeepAlive(context.TODO(), response.ID)
	if nil != err {
		return err
	}
	c.keepAliveChan = aliveRes
	return nil
}

//监测是否续约
func (c *ClientDis) listenLease() {
	for {
		select {
		case res := <-c.keepAliveChan:
			if nil == res {
				logrus.Error("租期续约失败，请查看Etcd日志")
				return
			}
			break
		}
	}
}

//获取节点数据
func (this *ClientDis) GetNodesInfo(path string) {
	resp, err := this.client.Get(context.TODO(), path, clientv3.WithPrefix())
	if err != nil {
		logrus.Error("Etcd 获取内容失败")
		return
	}
	//连接服务
	for key, value := range this.extractAddrs(resp) {
		//排除自己
		if key == this.path {
			continue
		}
		//不需要其他连接
		var _conf config.ServiceModel
		json.Unmarshal([]byte(value), &_conf)
		//判断对方是否需要被连接，如果不需要，则剔除
		if !_conf.IsConnect {
			continue
		}
		//连接grpc服务
		client := grpc_service.ConnectGRPCService([]byte(value))
		if client == nil { //连接不上的时候，则不再连接，从服务中剔除
			continue
		}
		//添加已连接的服务列表中
		EtcdClient.ConnServiceList[key] = &EtcdConnStruct{
			ConfigData:  value,
			ConnectName: key,
			Connect:     client,
		}
	}
}

//监听节点数据变化
func (this *ClientDis) WatchNodes(key string) {
	watcher := clientv3.NewWatcher(this.client)
	path := this.path + "/" + key
	for {
		rch := watcher.Watch(context.TODO(), path, clientv3.WithPrefix())
		for wresp := range rch {
			for _, event := range wresp.Events {
				switch event.Type {
				case mvccpb.PUT:
					//todo 维护一个列表
					var _conf config.ServiceModel
					json.Unmarshal(event.Kv.Value, &_conf)
					if !_conf.IsConnect {
						continue
					}
					if nil != this.ConnServiceList[string(event.Kv.Key)] {
						//如果已连接节点，无须再连接
						continue
					}
					client := grpc_service.ConnectGRPCService(event.Kv.Value)
					if nil == client {
						continue
					}
					this.ConnServiceList[string(event.Kv.Key)] = &EtcdConnStruct{
						ConfigData:  string(event.Kv.Value),
						ConnectName: string(event.Kv.Key),
						Connect:     client,
					}
				case mvccpb.DELETE:
					//TODO 删除etcd剔除的服务，首先从服务器断掉该连接，然后再删除该数据
					delete(this.ConnServiceList, string(event.Kv.Key))
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
func (this *ClientDis) DelNode(key string) {
	this.Lock()
	defer this.Unlock()
	path := this.path + "/" + key
	response, err := this.client.Delete(context.TODO(), path, clientv3.WithPrefix())
	if nil != err {
		logrus.Println("Error:", err.Error())
	}
	logrus.Println("Delete node", response.Deleted)
}
