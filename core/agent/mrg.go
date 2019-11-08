package agent

import (
	"YaIce/core/cluster"
	"YaIce/core/config"
	"YaIce/core/yaml"
	"context"
	"encoding/json"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
)

const TTL = 20

var mrg *AgentModel

type AgentModel struct {
	sync.RWMutex
	context.Context
	Client        *clientv3.Client
	Conns         []string                     //需要连接服务列表
	Key           string                       //服务在Etcd中的key
	LeaseRes      *clientv3.LeaseGrantResponse //自己配置租约
	KeepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
}

func Init() {
	//连接的服务
	conns := strings.Split(yaml.YamlDevMrg.EtcdConnectString, "#")
	//唯一key值
	key := yaml.YamlDevMrg.ClusterName + "/" + config.Config.GroupName + "/" + config.Config.TypeName + "/" + config.Config.Pid
	//初始化代理、连接
	mrg = &AgentModel{Conns: conns, Key: key}
	if nil != Connect() {
		mrg = nil
	}
}

//连接etcd服务
func Connect() error {
	//连接etcd服务
	etcdCli, err := clientv3.New(clientv3.Config{Endpoints: mrg.Conns, DialTimeout: 5 * time.Second})
	if nil != err {
		logrus.Debug("Etcd 服务，启动错误，Error Msg：", err.Error())
		return err
	}
	//赋值etcdCli
	mrg.Client = etcdCli
	//监听一个服务器下的节点数据变化
	go mrg.watchNodes()
	return nil
}

func RegisterData() error {
	if nil == mrg {
		zap.Error(errors.New("agent not initiated"))
		return errors.New("agent not initiated")
	}
	mrg.Lock()
	defer mrg.Unlock()
	data, jsonErr := json.Marshal(config.Config)
	if nil != jsonErr {
		zap.Error(errors.New("json Serialization failed"))
		return errors.New("json Serialization failed")
	}
	//保持连接的时间
	keepErr := mrg.grantSetLeaseKeepAlive(TTL)
	if nil != keepErr {
		zap.Error(keepErr)
		return keepErr
	}
	//etcd中存储格式是KV
	_, err := mrg.Client.Put(context.TODO(), mrg.Key, string(data), clientv3.WithLease(mrg.LeaseRes.ID))
	if err != nil {
		return err
	}
	go mrg.listenLease()
	return nil
}

//获取节点数据
func GetNodeData(key string) map[string]*config.ModuleConfig {
	serviceConfList := make(map[string]*config.ModuleConfig)

	resp, err := mrg.Client.Get(context.TODO(), yaml.YamlDevMrg.ClusterName+"/"+key, clientv3.WithPrefix())
	if err != nil {
		logrus.Error("Etcd 获取内容失败")
		return serviceConfList
	}
	for key, value := range mrg.readNodeData(resp) {
		var _conf config.ModuleConfig
		json.Unmarshal([]byte(value), &_conf)
		serviceConfList[key] = &_conf
	}
	return serviceConfList
}

//删除节点
func DelNode(key string) {
	mrg.Lock()
	defer mrg.Unlock()
	response, err := mrg.Client.Delete(context.TODO(), key)
	if nil != err {
		logrus.Println("Error:", err.Error())
	}
	logrus.Println("Delete node", response.Deleted)
}

//读取节点数据
func (this *AgentModel) readNodeData(resp *clientv3.GetResponse) map[string]string {
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
func (this *AgentModel) grantSetLeaseKeepAlive(ttl int64) error {
	response, err := this.Client.Lease.Grant(context.TODO(), ttl)
	if nil != err {
		return err
	}
	this.LeaseRes = response
	aliveRes, err := this.Client.KeepAlive(context.TODO(), response.ID)
	if nil != err {
		return err
	}
	this.KeepAliveChan = aliveRes
	return nil
}

//监测是否续约
func (this *AgentModel) listenLease() {
	for {
		select {
		case res := <-this.KeepAliveChan:
			if nil == res {
				RegisterData()
				return
			}
			break
		}
	}
}

//监听节点数据变化
func (this *AgentModel) watchNodes() {
	watcher := clientv3.NewWatcher(this.Client)
	for {
		rch := watcher.Watch(context.TODO(), yaml.YamlDevMrg.ClusterName, clientv3.WithPrefix())
		for response := range rch {
			for _, event := range response.Events {
				key := string(event.Kv.Key)
				data := &config.ModuleConfig{}
				json.Unmarshal(event.Kv.Value, data)
				switch event.Type {
				case mvccpb.PUT:
					//判断新服务，是否是自己所需要的
					for _, conn := range config.Config.ConnServerNameList {
						if data.TypeName == conn {
							//排除属于自己的额
							if this.Key == key {
								continue
							}
							//服务是否已经在连接池中
							if data.GroupName == "center" || data.GroupName == config.Config.GroupName {
								//连接该服务
								if conn := cluster.Connect(data); nil != conn {
									//连接列表 集群名称+服务器组编号+服务类型
									logrus.Println("connect cluster error ...")
									continue
								}
							}
						}
					}
				case mvccpb.DELETE:
					DelNode(string(event.Kv.Key))
					//TODO 删除etcd剔除的服务，首先从服务器断掉该连接，然后再删除该数据
					cluster.Delete(key)
				}
			}
		}
	}
}
