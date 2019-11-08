package network

import (
	"YaIce/core/common"
	"YaIce/core/config"
	"YaIce/core/model"
	"YaIce/core/router"
	"YaIce/core/yaml"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	"io"
	"strconv"
)

//初始化外网监听
func Listen() error {
	for port := yaml.YamlDevMrg.NetworkPortStart; port <= yaml.YamlDevMrg.NetworkPortEnd; port++ {
		_port := ListenAccpet(port)
		if -1 != _port {
			config.Config.OutPort = strconv.Itoa(_port)
			return nil
		}
	}
	return errors.New("没有监听的端口")
}

//监听端口(kcp)
func ListenAccpet(port int) int {
	kcpListen, err := kcp.ListenWithOptions(":"+strconv.Itoa(port), nil, 10, 3)
	if nil != err {
		kcpConnsPtr.KcpListen = nil
		return -1
	}
	kcpConnsPtr.KcpListen = kcpListen
	go func() {
		for {
			conn, err := kcpListen.AcceptKCP()
			if nil != err || nil == conn {
				continue
			}
			kcpConnsPtr.MutexConns.Lock()
			if len(kcpConnsPtr.ConnectList) >= kcpConnsPtr.MaxConnect {
				//todo  返回客户端 服务器承载已满
				continue
			}
			//从conn读取玩家的playerGuid
			if nil == kcpConnsPtr.ConnectList[conn] {
				kcpConnsPtr.ConnectList[conn] = model.InitConn(conn)
			}
			kcpConnsPtr.MutexConns.Unlock()
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
		kcpConnsPtr.ChanMsgReviceQueue <- &model.MsgQueue{
			MsgNumber: protoNum,
			Session:   kcpConnsPtr.ConnectList[conn],
			MsgData:   buffer[4:n],
		}
	}
}

func Run() {
	if nil == kcpConnsPtr.KcpListen {
		return
	} else {
		for {
			select {
			case msg := <-kcpConnsPtr.ChanMsgReviceQueue:
				//需要增加过滤器
				router.CallFilterHandler(msg.Session, msg.MsgData)
				router.CallExternalRouterHandler(msg.MsgNumber, msg.Session, msg.MsgData)
				break
			case msg := <-kcpConnsPtr.ChanMsgSendQueue:
				msg.Session.WriteMsg(msg.MsgData)
				break
			default:
				break
			}
		}
	}
}

func SendMsg(session *model.Conn, protoData proto.Message) {
	protoNum := common.ProtocalNumber(common.GetProtoName(protoData))
	data, _ := proto.Marshal(protoData)
	content := common.IntToBytes(protoNum)
	content = append(content, data...)
	kcpConnsPtr.ChanMsgSendQueue <- &model.MsgQueue{
		Session:   session,
		MsgNumber: protoNum,
		MsgData:   content,
	}
}
