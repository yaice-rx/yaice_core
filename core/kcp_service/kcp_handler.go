package kcp_service

import (
	"YaIce/core/common"
	"YaIce/core/model"
	"YaIce/core/router"
	"YaIce/core/temp"
	"fmt"
	"github.com/xtaci/kcp-go"
	"io"
	"strconv"
)

//初始化外网监听
func ServerExternalInit()int{
	for port := temp.ConfigCacheData.YamlConfigData.PortStart; port <= temp.ConfigCacheData.YamlConfigData.PortEnd; port++{
		_port :=  serviceListenAccpet(port)
		if -1 != _port{
			return _port
		}
	}
	return -1
}

//监听端口(kcp)
func serviceListenAccpet(port int)int{
	kcpListen, err 	:= kcp.ListenWithOptions(":"+strconv.Itoa(port), nil, 10, 1)
	if nil != err{
		return -1
	}
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
			if len(KcpConnPtr.ConnectList) >= KcpConnPtr.MaxConnect{
				//todo  返回客户端 服务器承载已满
				continue
			}
			//todo 从在线cache用户中取值
			if nil == KcpConnPtr.ConnectList[conn]{
				//todo 从登陆服务器取值，获取该用户已经登陆
				KcpConnPtr.mutexConns.Lock()
				_conn := model.InitPlayerConn(conn)
				KcpConnPtr.ConnectList[conn] = _conn
				KcpConnPtr.mutexConns.Unlock()
			}
			//分配请求句柄
			go handleKcpMux(conn)
		}
	}()
	return port
}

//（kcp）处理数据
func handleKcpMux(conn *kcp.UDPSession) {
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
		if KcpConnPtr.ConnectList[conn] != nil {
			protoNum := common.BytesToInt(buffer[:4])
			//检测除登陆接口，其余全部检测合法性
			router.RouterListPtr.CallRouterHandler(protoNum,KcpConnPtr.ConnectList[conn],buffer[4:n])
		}
	}
}
