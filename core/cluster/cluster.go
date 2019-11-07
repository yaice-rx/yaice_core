package cluster

import (
	"YaIce/core/config"
	"YaIce/core/network"
	"YaIce/core/yaml"
	"errors"
	"github.com/xtaci/kcp-go"
	"strconv"
	"strings"
)

//example map[服务集群编号]map[服务类型]map[服务编号][]所连接的服务句柄
type _cluster struct {
	config  *config.ModuleConfig
	session *kcp.UDPSession
}

type clusterConns struct {
	ConnList map[string]map[string]map[string][]*_cluster
}

var ClusterMrg *clusterConns

//启动连接GRPCService服务
func Connect(data *config.ModuleConfig) error {
	conn, err := kcp.DialWithOptions(data.InHost+":"+data.InPort, nil, 10, 1)
	if nil != err {
		return err
	}
	ClusterMrg.ConnList[yaml.YamlDevMrg.ClusterName][data.GroupName][data.TypeName] =
		append(ClusterMrg.ConnList[yaml.YamlDevMrg.ClusterName][data.GroupName][data.TypeName], &_cluster{
			config:  data,
			session: conn,
		})
	return nil
}

//服务器监听
func Init() error {
	ClusterMrg = new(clusterConns)
	ClusterMrg.ConnList = make(map[string]map[string]map[string][]*_cluster)
	for port := yaml.YamlDevMrg.NetworkPortStart; port <= yaml.YamlDevMrg.NetworkPortEnd; port++ {
		_port := network.ListenAccpet(port)
		if -1 != _port {
			config.Config.InPort = strconv.Itoa(_port)
			return nil
		}
	}
	return errors.New("没有监听的端口")
}

func Delete(key string) {
	keyMap := strings.Split(key, "/")
	for i := 0; i < len(ClusterMrg.ConnList[keyMap[0]][keyMap[1]][keyMap[2]]); i++ {
		if ClusterMrg.ConnList[keyMap[0]][keyMap[1]][keyMap[2]][i].config.Pid == keyMap[3] {
			ClusterMrg.ConnList[keyMap[0]][keyMap[1]][keyMap[2]] =
				append(ClusterMrg.ConnList[keyMap[0]][keyMap[1]][keyMap[2]][:i],
					ClusterMrg.ConnList[keyMap[0]][keyMap[1]][keyMap[2]][i:]...)
		}
	}
}
