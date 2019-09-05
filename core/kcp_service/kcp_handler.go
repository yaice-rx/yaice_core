package kcp_service

import (
	"YaIce/core/common"
	"YaIce/core/config"
	"YaIce/core/model"
	"YaIce/core/router"
	"github.com/xtaci/kcp-go"
	"io"
	"strconv"
)

//初始化外网监听
func Init() int {
	for port := config.GetYamlData().PortStart; port <= config.GetYamlData().PortEnd; port++ {
		_port := serviceListenAccpet(port)
		if -1 != _port {
			return _port
		}
	}
	return -1
}

//监听端口(kcp)
func serviceListenAccpet(port int) int {
	kcpListen, err := kcp.ListenWithOptions(":"+strconv.Itoa(port), nil, 10, 3)
	if nil != err {
		return -1
	}
	go func() {
		for {
			conn, err := kcpListen.AcceptKCP()
			if nil != err || nil == conn {
				continue
			}
			KcpConnPtr.MutexConns.Lock()
			if len(KcpConnPtr.ConnectList) >= KcpConnPtr.MaxConnect {
				//todo  返回客户端 服务器承载已满
				continue
			}
			//从conn读取玩家的playerGuid
			if nil == KcpConnPtr.ConnectList[conn] {
				KcpConnPtr.ConnectList[conn] = model.InitPlayerConn(conn)
			}
			KcpConnPtr.MutexConns.Unlock()
			// 开启协程处理客户端请求，防止一条请求未处理完时，另一条请求阻塞
			go handleKcpMux(conn)
		}
	}()
	return port
}

//（kcp）处理数据
func handleKcpMux(conn *kcp.UDPSession) {
	for {
		//read
		var buffer = make([]byte, 1024)
		n, e := conn.Read(buffer)
		if e != nil {
			if e == io.EOF {
				break
			}
			break
		}
		protoNum := common.BytesToInt(buffer[:4])

		KcpConnPtr.ChanMsgQueue <- &MsgQueue{
			msgNumber: protoNum,
			Session:   KcpConnPtr.ConnectList[conn],
			msgData:   buffer[4:n],
		}
	}
}

func Run() {
	for {
		select {
		case msg := <-KcpConnPtr.ChanMsgQueue:
			//需要增加过滤器
			router.CallFilterHandler(msg.Session, msg.msgData)
			router.CallExternalRouterHandler(msg.msgNumber, msg.Session, msg.msgData)
			break
		default:
			break
		}
	}
}
