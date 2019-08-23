package config

//服务配置

type ServiceModel struct {
	name        string //服务器名称
	groupId     string //服务器组编号
	serverExtra string
	outHost     string //内部连接ip
	outPort     int    //内部连接端口
	inHost      string //外部连接ip
	inPort      int    //内部连接端口
	IsConnect   bool   //是否需要连接
}

//服务配置处理
var confServiceSystem *ServiceModel

func InitServiceConf(name string, groupId string, serverExtra string, inHost string, outHost string, isConn bool) {
	confServiceSystem = &ServiceModel{
		name:        name,
		groupId:     groupId,
		serverExtra: serverExtra,
		inHost:      inHost,
		outHost:     outHost,
		IsConnect:   isConn,
	}
}

func GetName() string {
	return confServiceSystem.name
}

func GetGroupId() string {
	return confServiceSystem.groupId
}

func GetServerExtra() string {
	return confServiceSystem.serverExtra
}

func GetOutHost() string {
	return confServiceSystem.outHost
}

func GetOutPort() int {
	return confServiceSystem.outPort
}

func GetInHost() string {
	return confServiceSystem.inHost
}

func GetInPort() int {
	return confServiceSystem.inPort
}

func GetIsConn() bool {
	return confServiceSystem.IsConnect
}

func SetInPort(port int) {
	confServiceSystem.inPort = port
}

func SetOutPort(port int) {
	confServiceSystem.outPort = port
}

func GetServiceConfData() ServiceModel {
	if nil == confServiceSystem {
		return ServiceModel{}
	}
	return *confServiceSystem
}
