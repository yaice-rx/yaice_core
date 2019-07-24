package config

type ServiceConfig struct {
	ServerType 		string		//服务器类型
	ServerGroupId	string		//服务器组编号
	ConfigPathName 	string     	//配置文件路径
	InternalHost 	string     	//内部连接ip
	ExternalHost	string     	//外部连接ip
}

var ServiceConfigData ServiceConfig

func InitServiceConfig(_type string,groupId string,path string,i_host string,e_host string){
	ServiceConfigData = ServiceConfig{
		ServerType:_type,
		ServerGroupId:groupId,
		ConfigPathName:path,
		InternalHost:i_host,
		ExternalHost:e_host,
	}
}