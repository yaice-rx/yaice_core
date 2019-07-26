package config

type ServiceConfig struct {
	ServerName 		string		//服务器名称
	ServerGroupId	string		//服务器组编号
	ConfigPathName 	string     	//配置文件路径
	InternalHost 	string     	//内部连接ip
	InternalPort 	int     	//内部连接端口
	ExternalHost	string     	//外部连接ip
	ExternalPort 	int     	//内部连接端口
	IsConnect		bool		//是否需要连接
}

var ServiceConfigData ServiceConfig

func InitServiceConfig(name string,groupId string,path string,i_host string,e_host string,is_connect bool){
	ServiceConfigData = ServiceConfig{
		ServerName		:name,
		ServerGroupId	:groupId,
		ConfigPathName	:path,
		InternalHost	:i_host,
		ExternalHost	:e_host,
		IsConnect		:is_connect,
	}
}