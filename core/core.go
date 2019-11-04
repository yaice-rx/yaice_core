package core

import (
	"YaIce/core/agent"
	"YaIce/core/cluster"
	"YaIce/core/config"
	"YaIce/core/database"
	"YaIce/core/job"
	"YaIce/core/network"
	"YaIce/core/router"
)

var ConnCount int = 5000

type ModuleCore interface {
	//注册路由
	RegisterRouter()
	//监听端口或者连接
	Listen()
	//初始化配置
	StartHook()
	//监听http
	ListenHttp()
}

type Module struct {
	ModuleCore
}

func Run(m ModuleCore) {
	//初始化程序
	onInit()
	//启动程序所需配置
	m.StartHook()
	//连接网络
	m.Listen()
	//注册路由
	m.RegisterRouter()
	//执行程序
	onExec()
	//执行http监听端口
	m.ListenHttp()
}

func onInit() {
	//初始化网络连接信息
	network.Init(ConnCount)
	//初始化Etcd
	config.InitImplEtcd()
	//初始化路由
	router.Init()
	//初始化Etcd
	agent.Init()
	//初始化集群
	cluster.Init()
	//初始化数据库
	database.Connect()
}

func onExec() {
	//连接集群
	connClusterServer()
	//监听集群事件
	cluster.Listen()
	//加入集群
	agent.RegisterData()
	//初始化定时器
	job.Start()
	//启动服务
	network.Run()
}

//grpc连接
func connClusterServer() {
	for _, value := range agent.GetNodeData("") {
		//判断当前那些服务是自己需要连接的
		for _, self := range config.StartupConfigMrg.ConnServerNameList {
			if self == value.TypeName {
				//如果是中心服务 或者 属于自己分组内部服务 都是可以连接的
				if value.GroupName == "center" || value.GroupName == config.StartupConfigMrg.GroupName {
					cluster.Handler.Connect(config.ConfDevMrg.ClusterName+"/"+value.GroupName+"/"+value.TypeName+"/"+value.Pid, value.InHost+":"+value.InPort)
				}
			}
		}
	}
}
