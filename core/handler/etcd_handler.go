package handler

import (
	"YaIce/core/config"
	"YaIce/core/model"
	"context"
	"encoding/json"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"time"
)

const TTL = 20

var etcdCli *model.ClientModel

//连接etcd服务
func EtcdConnect(groupId string, serviceName string, etcdConn string) error {
	pid, err := uuid.NewV4()
	if err != nil {
		return errors.New("uuid must exist")
	}
	serverList := []string{etcdConn}
	//连接etcd服务
	clientCli, err := clientv3.New(clientv3.Config{Endpoints: serverList, DialTimeout: 5 * time.Second})
	if nil != err {
		logrus.Debug("Etcd 服务，启动错误，Error Msg：", err.Error())
		return err
	}
	//初始化连接信息
	EtcdInit(serverList, clientCli, serviceName, groupId+"/"+serviceName+"/"+pid.String())
	return nil
}

func RegisterEtcdData(data string) {
	if nil == etcdCli {
		return
	}
	grantSetLeaseKeepAlive(TTL)
	etcdCli.Lock()
	_, err := etcdCli.Client.Put(context.TODO(), etcdCli.Path, data, clientv3.WithLease(etcdCli.LeaseRes.ID))
	if err != nil {
		logrus.Debug("数据注册失败，Error Msg：", err.Error())
		return
	}
	etcdCli.Unlock()
	go listenLease()
}

//获取节点数据
func GetEtcdNodeData(path string) []config.ServiceConfigModel {
	serviceConfList := []config.ServiceConfigModel{}
	resp, err := etcdCli.Client.Get(context.TODO(), path, clientv3.WithPrefix())
	if err != nil {
		logrus.Error("Etcd 获取内容失败")
		return serviceConfList
	}
	for _, value := range readNodeData(resp) {
		var _conf config.ServiceConfigModel
		json.Unmarshal([]byte(value), &_conf)
		serviceConfList = append(serviceConfList, _conf)
	}
	return serviceConfList
}

//删除节点
func DelNode(key string) {
	etcdCli.Lock()
	defer etcdCli.Unlock()
	path := etcdCli.Path + "/" + key
	response, err := etcdCli.Client.Delete(context.TODO(), path, clientv3.WithPrefix())
	if nil != err {
		logrus.Println("Error:", err.Error())
	}
	logrus.Println("Delete node", response.Deleted)
}

func EtcdInit(etcdList []string, cli *clientv3.Client, name string, path string) {
	etcdCli = &model.ClientModel{
		Endpoints:   etcdList,
		Client:      cli,
		ServiceName: name,
		Path:        path,
	}
}

//读取节点数据
func readNodeData(resp *clientv3.GetResponse) map[string]string {
	data := make(map[string]string, 0)
	if resp == nil || resp.Kvs == nil {
		return data
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			key := string(resp.Kvs[i].Key)
			data[key] = string(v)
		}
	}
	return data
}

//授权租期，自动续约
func grantSetLeaseKeepAlive(ttl int64) error {
	response, err := etcdCli.Client.Lease.Grant(context.TODO(), ttl)
	if nil != err {
		return err
	}
	etcdCli.LeaseRes = response
	aliveRes, err := etcdCli.Client.KeepAlive(context.TODO(), response.ID)
	if nil != err {
		return err
	}
	etcdCli.KeepAliveChan = aliveRes
	return nil
}

//监测是否续约
func listenLease() {
	for {
		select {
		case res := <-etcdCli.KeepAliveChan:
			if nil == res {
				logrus.Error("租期续约失败，请查看Etcd日志")
				return
			}
			break
		}
	}
}

//监听节点数据变化
func watchNodes(key string) {
	watcher := clientv3.NewWatcher(etcdCli.Client)
	path := etcdCli.Path + "/" + key
	for {
		rch := watcher.Watch(context.TODO(), path, clientv3.WithPrefix())
		for wresp := range rch {
			for _, event := range wresp.Events {
				var _conf config.ServiceConfigModel
				json.Unmarshal(event.Kv.Value, &_conf)
				switch event.Type {
				case mvccpb.PUT:
					//如果没有所需要连接的服务
					if len(_conf.GetConnList()) <= 0 {
						continue
					}
					if nil != ServerMapHandler[_conf.GetName()][_conf.GetPid()] {
						//如果已连接节点，无须再连接
						return
					}
					GRPCConnect(ServerMapHandler[_conf.GetName()], _conf)
				case mvccpb.DELETE:
					//TODO 删除etcd剔除的服务，首先从服务器断掉该连接，然后再删除该数据
					DeleteGRPCConn(_conf.GetName(), _conf.GetPid())
				}
			}
		}
	}
}
