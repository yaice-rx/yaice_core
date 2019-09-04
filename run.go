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
	serverName := flag.String("name", "auth", "服务名称")
	//内网地址
	internal_host := flag.String("internal_host", "127.0.0.1", "对内监听地址")
	//外网地址
	external_host := flag.String("external_host", "", "对外监听地址")
	//外网监听http端口
	http_port := flag.String("http_port", "8080", "http监听端口")
	//服务分组Id（自己本身的服务）
	server_group := flag.String("server_group", "common", "服务分组")
	//需要连接的扩展服务
	server_extra := flag.String("server_extra", "common", "中心服务")
	//是否需要被其他服务连接
	is_connect := flag.Bool("is_connnect", false, "是否需要其他服务连接")
	//解析数据
	flag.Parse()
	//初始化服务配置
	config.InitServiceConf(*serverName, *server_group, *server_extra, *internal_host, *external_host, *is_connect)
	//初始化YAML数据
	config.YamlInit()
	//配置
	core.NewServerCore()
	//初始化调用对应的服务
	switch *serverName {
	case "auth":
		auth.Initialize(*http_port, *server_group)
		return
	case "game":
		game.Initialize()
		return
	case "client":
		client.Initialize()
		return
	}
}
