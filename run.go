package main

import (
	"YaIce/client"
	"YaIce/core"
	"YaIce/game"
	"YaIce/gate"
	"flag"
)

//start server
func main() {
	//服务器类型
	serverType := flag.String("type", "gate", "a string var")
	//模块配置配置
	appConfigName := flag.String("appConfigName", "", "app config name")
	//内网地址
	internal_host := flag.String("internal_host", "127.0.0.1", "internal host")
	//外网地址
	external_host := flag.String("external_host", "", "external host")
	//外网监听http端口
	http_port := flag.String("http_port", "8080", "host port")

	server_id := flag.String("server_id","-1","server_id")

	flag.Parse()
	//配置
	core := core.NewServerCore()
	core.ConfigFileName = *appConfigName
	core.ExternalHost 	= *internal_host
	core.InternalHost 	= *external_host
	core.ServerType		= *serverType
	//初始化调用对应的服务
	switch *serverType {
		case "gate":
			gate.Initialize(core,*http_port,*server_id)
			return
		case "game":
			game.Initialize(core,*server_id)
			return
		case "client":
			client.Initialize(core)
			return
	}
}
