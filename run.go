package main

import (
	"YaIce/auth"
	"YaIce/client"
	"YaIce/core"
	"YaIce/core/config"
	"YaIce/game"
	"flag"
)

//start server
func main() {
	//服务器名称
	serverType := flag.String("type", "auth", "服务类型")
	//内网地址
	in_host := flag.String("in_host", "127.0.0.1", "对内监听地址")
	//外网地址
	outer_host := flag.String("outer_host", "", "对外监听地址")
	//http端口
	http_port := flag.String("http_port", "8080", "http监听端口")
	//服务分组Id [center代表服务列表中中心处理器]
	server_group := flag.String("server_group", "center", "服务分组")
	//需要连接的扩展服务
	connect_service_list := flag.String("connect_service_list", "", "需要连接的服务类型")
	//解析数据
	flag.Parse()
	//初始化服务配置
	config.Init(*serverType, *server_group, *connect_service_list, *in_host, *outer_host, *http_port)
	//初始化调用对应的服务
	switch *serverType {
	case "auth":
		core.Run(auth.ModuleMrg)
		return
	case "game":
		core.Run(game.ModuleMrg)
		return
	case "client":
		client.Initialize()
		return
	case "close":
		return
	}
}
