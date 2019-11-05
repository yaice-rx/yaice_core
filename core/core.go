package core

import (
	"YaIce/core/agent"
	"YaIce/core/cluster"
	"YaIce/core/config"
	"YaIce/core/database"
	"YaIce/core/job"
	"YaIce/core/network"
	"YaIce/core/router"
	"YaIce/core/yaml"
)

var ConnCount int = 5000

type ModuleCore interface {
	//注册钩子
	RegisterHook()
	//注册路由
	RegisterRouter()
	//监听端口或者连接
	Listen()
}

type Module struct {
	ModuleCore
}

var ModuleMrg *Module

func onInit() {
	ModuleMrg = new(Module)
	network.Init(ConnCount)
	yaml.Init()
	router.Init()   //router 注册
	agent.Init()    //etcd 注册
	cluster.Init()  //grpc注册
	database.Init() //数据库注册
	job.Init()      //定时器注册
}

func Run(m ModuleCore) {
	onInit()
	m.RegisterHook()
	m.RegisterRouter()
	connClusterServer()
	cluster.Listen()
	agent.RegisterData()
	m.Listen()
}

//grpc连接
func connClusterServer() {
	for _, value := range agent.GetNodeData("") {
		//判断当前那些服务是自己需要连接的
		for _, self := range config.Config.ConnServerNameList {
			if self == value.TypeName {
				//如果是中心服务 或者 属于自己分组内部服务 都是可以连接的
				if value.GroupName == "center" ||
					value.GroupName == config.Config.GroupName {
					cluster.Handler.Connect(
						yaml.YamlDevMrg.ClusterName+"/"+value.GroupName+"/"+value.TypeName+"/"+value.Pid,
						value.InHost+":"+value.InPort)
				}
			}
		}
	}
}
