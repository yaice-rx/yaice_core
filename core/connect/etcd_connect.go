package connect

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

type ClientDis struct {
	sync.RWMutex
	client 		*clientv3.Client
	Endpoints   []string
	serverId 	int
	serverType	string
	leaseRes    *clientv3.LeaseGrantResponse
	keepAliveChan  <-chan *clientv3.LeaseKeepAliveResponse
}

const ttl  = 200

//初始化Etcd服务，
//存储结构表如下：
//			serverId:1=>{"gate_序号":地址，"game_序号"：地址},
func InitEtcd(serverId int,serverType string)(*ClientDis,error){
	etcdServerList :=  []string{"localhost:2379"}
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdServerList,
		DialTimeout: 5 * time.Second,
	})
	if nil != err {
		logrus.Debug(err.Error())
		return nil,err
	}
	return &ClientDis{
		Endpoints:etcdServerList,
		client:etcdCli,
		serverId:serverId,
		serverType:serverType,
	},nil
}

func (c *ClientDis)RegisterNode(key string ,value string){
	if nil == c {
		return
	}
	c.grantSetLeaseKeepAlive(ttl)
	c.Lock()
	path := strconv.Itoa(c.serverId)+"/"+c.serverType+"/"+key
	putResp, err := c.client.Put(context.TODO(),path,value,clientv3.WithLease(c.leaseRes.ID));
	if err != nil{
		fmt.Println(err)
		return
	}
	c.Unlock()
	go c.listenLease()
	fmt.Println("写入成功:", putResp.Header.Revision)
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

//监听
func (c *ClientDis)listenLease(){
	for  {
		select {
		case res :=  <- c.keepAliveChan:
			if nil == res{
				fmt.Println("lease close")
				return
			}
			fmt.Println("lease success")
			break;
		}
	}
}

//获取节点数据
func (c *ClientDis)GetNodesInfo(key string)([]string,error){
	path := strconv.Itoa(c.serverId)+"/"+c.serverType+"/"+key
	resp, err := c.client.Get(context.TODO(),path)
	if err != nil {
		return nil,err
	}
	return c.extractAddrs(resp),nil
}

//监听节点数据变化
func (this *ClientDis) WatchNodes(key string){
	watcher := clientv3.NewWatcher(this.client)
	path := strconv.Itoa(this.serverId)+"/"+this.serverType
	for {
		rch := watcher.Watch(context.TODO(), path, clientv3.WithPrefix())
		for wresp := range rch {
			for _, event := range wresp.Events {
				switch (event.Type) {
				case mvccpb.PUT:
					fmt.Println("PUT事件",event.Kv.Key,event.Kv.Value)
				case mvccpb.DELETE:
					fmt.Println("DELETE事件",event.Kv.Key,event.Kv.Value)
				}
			}
		}
	}
}

//读取节点数据
func (this *ClientDis) extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}

func (this *ClientDis)DelNode(key string)  {
	this.Lock()
	defer this.Unlock()
	path := strconv.Itoa(this.serverId)+"/"+this.serverType+"/"+key
	response,err := this.client.Delete(context.TODO(),path,clientv3.WithPrefix())
	if nil != err{
		logrus.Println(err.Error())
	}
	logrus.Println("Delete node",response.Deleted)

}