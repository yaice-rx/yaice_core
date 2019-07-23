package connect

import (
	"YaIce/core"
	"YaIce/core/etcd_service"
	"YaIce/core/grpc_service"
	"YaIce/core/temp"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

//开启grpc服务模式
func ServerInternalInit()int{
	//从zookeeper中获取登陆服务器的ip
	server := grpc.NewServer()
	//注册路由
	grpc_service.RegisterServiceGrpc(server)
	reflection.Register(server)
	//获取 端口
	for port := temp.ConfigCacheData.YamlConfigData.PortStart; port <= temp.ConfigCacheData.YamlConfigData.PortEnd; port++{
		address, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err == nil {
			go server.Serve(address)
			return port;
		}
	}
	return -1
}

//连接网管服务器内部
func RegisterGateService(){
	//连接auth服务器
	ConnectService(core.ServerCoreHandler.ServerGroupId+"/auth")
}

func RegisterRelationService(){
	//连接Relation服务器
	ConnectService("/relation")
}

//连接服务器
func ConnectService(path string){
	jsonData,err :=  etcd_service.EtcdClient.GetNodesInfo(path)
	if nil != err{
		logrus.Debug(err.Error())
		return
	}
	for i := 0; i < len(jsonData);i++{
		var etcdData etcd_service.ServerConfigEtcd
		json.Unmarshal([]byte(jsonData[i]),&etcdData)
		conn, err := grpc.Dial(etcdData.InternalIP+":"+etcdData.InternalPort, grpc.WithInsecure())
		if nil != err{
			continue
		}
		grpc_service.RegisterClientGrpc(conn)
	}
}