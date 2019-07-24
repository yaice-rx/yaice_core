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
	//服务器类型
	serverType 		:= flag.String("type", "gate", "a string var")
	//模块配置配置
	configPath 		:= flag.String("config_path", "", "config_path")
	//内网地址
	internal_host 	:= flag.String("internal_host", "127.0.0.1", "internal host")
	//外网地址
	external_host 	:= flag.String("external_host", "", "external host")
	//外网监听http端口
	http_port 		:= flag.String("http_port", "8080", "host port")
	//服务器id
	server_group_id := flag.String("server_group_id","","server_group_id")
	//解析数据
	flag.Parse()
	//配置
	core.NewServerCore()

	config.InitServiceConfig(*serverType,*server_group_id,*configPath,*internal_host,*external_host)

	core.ServerCoreHandler.ServiceConfig = config.ServiceConfigData
	//初始化调用对应的服务
	switch *serverType {
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
