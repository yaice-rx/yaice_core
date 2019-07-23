package core

import (
	"YaIce/core/common"
	"YaIce/core/connect"
	"YaIce/core/job"
	"YaIce/core/model"
	"YaIce/core/router"
	"YaIce/core/temp"
	"fmt"
	"github.com/xtaci/kcp-go"
	"io"
	"strconv"
	"sync"
)

type ServerCore struct{
	Config          *temp.ConfigModule	//配置数据
	maxConnect		int					//最大连接数据
	mutexConns      sync.Mutex
	ServerType 		string
	ServerGroupId	string						//服务器组编号
	ConfigFileName 	string                      //配置文件名称
	InternalHost 	string                      //内部连接ip
	ExternalHost	string                      //外部连接ip
	Routers			*router.RegisterRouterRequest      //注册回调方法列表
	TickTasks		map[string]func()             //tick函数列表
	ConnectList 	map[*kcp.UDPSession]*connect.PlayerConn // uid->连接Conn
	DB 				*model.DBModel               //数据库
}

var ServerCoreHandler *ServerCore

func NewServerCore()  {
	s := new(ServerCore)
	//初始化路由
	s.Routers 		= new(router.RegisterRouterRequest)
	//初始化连接容量
	s.ConnectList	= make(map[*kcp.UDPSession]*connect.PlayerConn)
	//最大连接数
	s.maxConnect 	= 5000
	//初始化数据库连接
	s.DB 			= model.Init()
	//开启定时任务
	go job.CallJob()
	ServerCoreHandler = s
}

//初始化外网监听
func (s *ServerCore)ServerExternalInit()int{
	for port := temp.ConfigCacheData.YamlConfigData.PortStart; port <= temp.ConfigCacheData.YamlConfigData.PortEnd; port++{
		_port :=  s.ServerListenAccpet(port)
		if -1 != _port{
			return _port
		}
	}
	return -1
}

//监听端口(kcp)
func (s *ServerCore)ServerListenAccpet(port int)int{
	kcpListen, err 	:= kcp.ListenWithOptions(":"+strconv.Itoa(port), nil, 10, 1)
	if nil != err{
		return -1
	}
	//启动grpc
	go func(){
		for{
			conn, err := kcpListen.AcceptKCP()
			if nil != err{
				fmt.Println(err.Error())
				continue
			}
			if nil == conn{
				continue
			}
			if len(s.ConnectList) >= s.maxConnect{
				fmt.Println("too many connections")
				continue
			}
			//todo 从在线cache用户中取值
			if nil == s.ConnectList[conn]{
				//todo 从登陆服务器取值，获取该用户已经登陆
				s.mutexConns.Lock()
				_conn := connect.InitPlayerConn(conn)
				s.ConnectList[conn] = _conn
				s.mutexConns.Unlock()
			}
			//分配请求句柄
			go s.handleKcpMux(conn)
		}
	}()
	return port
}

 //（kcp）处理数据
func (s *ServerCore)handleKcpMux(conn *kcp.UDPSession) {
	var buffer = make([]byte,1024)
	for {
		n,e := conn.Read(buffer)
		if e != nil {
			if e == io.EOF{
				 break
			}
			break
		}
		//从conn读取玩家的playerGuid
		if s.ConnectList[conn] != nil {
			protoNum := common.BytesToInt(buffer[:4])
			//检测除登陆接口，其余全部检测合法性
			s.Routers.CallRouterHandler(protoNum,s.ConnectList[conn],buffer[4:n])
		}
	}
}






