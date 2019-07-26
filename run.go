package main

import (
	"YaIce/client"
	"YaIce/core"
	"YaIce/core/config"
	"YaIce/game"
	"YaIce/gate"
	"flag"
)

//start server
func main() {
	//服务器名称
	serverName 		:= flag.String("name", "gate", "a string var")
	//模块配置配置
	configPath 		:= flag.String("config_path", "", "config_path")
	//内网地址
	internal_host 	:= flag.String("internal_host", "127.0.0.1", "internal host")
	//外网地址
	external_host 	:= flag.String("external_host", "", "external host")
	//外网监听http端口
	http_port 		:= flag.String("http_port", "8080", "host port")
	//服务分组Id 如果为中心服务，则在启动参数中不添加,默认为common，中心服务器
	server_group_id := flag.String("server_group_id","common","server_group_id")
	//是否需要连接
	is_connect := flag.Bool("is_connnect",false,"is connect")
	//解析数据
	flag.Parse()
	//配置
	core.NewServerCore()

	config.InitServiceConfig(*serverName,*server_group_id,*configPath,*internal_host,*external_host,*is_connect)

	core.ServerCoreHandler.ServiceConfig = config.ServiceConfigData
	//初始化调用对应的服务
	switch *serverName {
		case "gate":
			gate.Initialize(*http_port,*server_group_id)
			return
		case "game":
			game.Initialize()
			return
		case "client":
			client.Initialize()
			return
	}
}
