package core

import (
	"YaIce/core/common"
	"YaIce/core/connect"
	"YaIce/core/job"
	"YaIce/core/model"
	"YaIce/core/temp"
	"fmt"
	"github.com/spf13/viper"
	"github.com/xtaci/kcp-go"
	"io"
	"runtime"
	"strconv"
	"sync"
)

type ServerCore struct{
	Config          *temp.ConfigModule	//配置数据
	maxConnect		int					//最大连接数据
	mutexConns      sync.Mutex
	ServerType 		string
	ConfigFileName 	string                      //配置文件名称
	InternalHost 	string                      //内部连接ip
	ExternalHost	string                      //外部连接ip
	Routers			*RegisterRouterRequest      //注册回调方法列表
	TickTasks		map[string]func()             //tick函数列表
	ConnectList 	map[*kcp.UDPSession]*connect.PlayerConn // uid->连接Conn
	DB 				*model.DBModel               //数据库
}

func NewServerCore() *ServerCore {
	s := new(ServerCore)
	//最大连接数
	s.maxConnect 	= 5000
	//初始化数据库连接
	s.DB 			= model.Init()
	//开启定时任务
	go job.CallJob()
	//检测客户端->服务器是否超时
	job.JoinJob(4,s.checkConnectTimeOut)
	return s
}

//初始化外网监听
func (s *ServerCore)ServerExternalInit(){
	listenPort :=  temp.ConfigCacheData.YamlConfigData.TcpServicePortStart
	kcpListen, err 	:= kcp.ListenWithOptions(":"+strconv.Itoa(), nil, 10, 1)
	defer func() {
		kcpListen.Close()
	}()
	if err != nil {
		panic(err.Error())
	}
	go func(){
		for{
			conn, err := kcpListen.AcceptKCP()
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
			if  err == nil {
				go s.handleMux(conn)
			} else {
				fmt.Println(err.Error())
			}

		}
	}()
	select {
	}
}

func (s *ServerCore)ServerInternalInit(){
	//从zookeeper中获取登陆服务器的ip

}
 //处理数据
func (s *ServerCore)handleMux(conn *kcp.UDPSession) {
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

//加载配置文件
func (s *ServerCore)loadConfig() string {
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("conf")
	//添加读取的配置文件路径
	v.AddConfigPath("./")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	var str string
	switch runtime.GOOS {
		case "darwin":
			str = "%GOPATH/src/YaIce"
			break
		case "windows":
			str = "%GOPATH/src/YaIce"
			break
		case "linux":
			str = "$GOPATH/src/YaIce"
			break
	}
	v.AddConfigPath(str)
	//设置配置文件类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig();err == nil {
		return v.GetString("ExcelPath")
	}
	return ""
}

//检查链接是否超时
func (s *ServerCore)checkConnectTimeOut(){
	/*for k,v := range s.ConnectList{
		if v.updateConnectTime + 5 < time.Now().Unix() {
			s.ConnectList[k] = nil
		}
	}*/
}



