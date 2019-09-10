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
	/*cpuProfile, _ := os.Create("cpu_profile_"+uuid.Must(uuid.NewV4()).String())
	pprof.StartCPUProfile(cpuProfile)*/
	//服务器名称
	serverName := flag.String("name", "auth", "服务名称")
	//内网地址
	in_host := flag.String("in_host", "127.0.0.1", "对内监听地址")
	//外网地址
	outer_host := flag.String("outer_host", "", "对外监听地址")
	//外网监听http端口
	http_port := flag.String("http_port", "8080", "http监听端口")
	//服务分组Id（自己本身的服务）
	server_group := flag.String("server_group", "1", "服务分组")
	//需要连接的扩展服务
	connect_service := flag.String("connect_service", "", "需要连接的服务")
	//解析数据
	flag.Parse()
	//初始化服务配置
	config.InitServiceConf(*serverName, *server_group, *connect_service, *in_host, *outer_host)
	//配置
	core.NewServerCore()
	//初始化定时器
	/*job.Crontab.AddCronTask(1, 200, func() {
		pprof.StopCPUProfile()
	})*/
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
