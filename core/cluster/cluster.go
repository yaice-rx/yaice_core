package cluster

import (
	"YaIce/core/config"
	"YaIce/core/model"
	"YaIce/core/network"
	"YaIce/core/proto"
	"YaIce/core/router"
	"YaIce/core/yaml"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	"strconv"
	"strings"
)

//连接服务的句柄 map[服务集群编号]map[服务类型]map[服务编号][pid]所连接的服务句柄
type clusterConns struct {
	ConnServiceList map[string]map[string]map[string]map[string]*model.Conn
	ConnClientList  map[string]map[string]map[string]map[string]*model.Conn
}

var ClusterMrg *clusterConns

//启动连接GRPCService服务
func Connect(data *config.ModuleConfig) error {
	conn, err := kcp.DialWithOptions(data.InHost+":"+data.InPort, nil, 10, 1)
	if nil != err {
		return err
	}

	defer func() {
		data := _proto.C2GGameRegister{
			GroupId:  config.Config.GroupName,
			TypeName: config.Config.TypeName,
			Pid:      config.Config.Pid}
		network.SendMsg(model.InitConn(conn), &data)
	}()
	if nil != ClusterMrg.ConnServiceList[yaml.YamlDevMrg.ClusterName] {
		ClusterMrg.ConnServiceList[yaml.YamlDevMrg.ClusterName][data.GroupName][data.TypeName][data.Pid] =
			model.InitConn(conn)
	} else {
		connMap := make(map[string]*model.Conn)
		connMap[data.Pid] = model.InitConn(conn)
		_typeConnMap := make(map[string]map[string]*model.Conn)
		_typeConnMap[data.TypeName] = connMap
		_groupConnMap := make(map[string]map[string]map[string]*model.Conn)
		_groupConnMap[data.GroupName] = _typeConnMap
		ClusterMrg.ConnServiceList[yaml.YamlDevMrg.ClusterName] = _groupConnMap
	}
	return nil
}

//服务器监听
func Init() error {
	defer func() {
		registerRouter()
		ClusterMrg = new(clusterConns)
		ClusterMrg.ConnServiceList = make(map[string]map[string]map[string]map[string]*model.Conn)
		ClusterMrg.ConnClientList = make(map[string]map[string]map[string]map[string]*model.Conn)
	}()
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
	delete(ClusterMrg.ConnServiceList[keyMap[0]][keyMap[1]][keyMap[2]], keyMap[3])
}

func registerRouter() {
	router.RegisterRouterHandler(&_proto.C2GGameRegister{}, serviceGameRegister)
}

func serviceGameRegister(conn *model.Conn, data []byte) {
	protoData := _proto.C2GGameRegister{}
	proto.Unmarshal(data, &protoData)
	if nil != ClusterMrg.ConnClientList[yaml.YamlDevMrg.ClusterName] {
		ClusterMrg.ConnClientList[yaml.YamlDevMrg.ClusterName][protoData.GroupId][protoData.TypeName][protoData.Pid] =
			conn
	} else {
		connMap := make(map[string]*model.Conn)
		connMap[protoData.Pid] = conn
		_typeConnMap := make(map[string]map[string]*model.Conn)
		_typeConnMap[protoData.TypeName] = connMap
		_groupConnMap := make(map[string]map[string]map[string]*model.Conn)
		_groupConnMap[protoData.GroupId] = _typeConnMap
		ClusterMrg.ConnClientList[yaml.YamlDevMrg.ClusterName] = _groupConnMap
	}
}
