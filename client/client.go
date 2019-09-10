package client

import (
	"YaIce/core/common"
	"YaIce/core/job"
	"YaIce/core/model"
	"YaIce/protobuf/external"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var conn *kcp.UDPSession
var token model.LoginToken

func Initialize() {
	resp, err := http.Post("http://10.0.0.10:8888/login",
		"application/x-www-form-urlencoded",
		strings.NewReader("userName=admin"))
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &token)
	kcpconn, err := kcp.DialWithOptions(token.Host+":"+strconv.Itoa(token.Port), nil, 10, 1)
	defer kcpconn.Close()
	if err != nil {
		fmt.Println("kcp err", err.Error())
		return
	}
	conn = kcpconn
	RegisterHandler()
	job.Crontab.AddCronTask(5, -1, pingHandler)
	go handleKcpConn(conn)
	select {}
}
func RegisterHandler() {
	gmCommand := c2game.C2GRegister{Pid: token.Pid, SessionId: token.SessionKey}
	data, err := proto.Marshal(&gmCommand)
	if err != nil {
		log.Fatalln("Marshal mrg data error: ", err)
	}
	SendMsg(conn, common.ProtocalNumber(common.GetProtoName(&c2game.C2GRegister{})), data)
}
func pingHandler() {
	gmCommand := c2game.C2GPing{}
	data, err := proto.Marshal(&gmCommand)
	if err != nil {
		log.Fatalln("Marshal mrg data error: ", err)
	}
	SendMsg(conn, common.ProtocalNumber(common.GetProtoName(&c2game.C2GPing{})), data)
}

func SendMsg(conn *kcp.UDPSession, protoNumber int32, data []byte) {
	if conn != nil {
		content := common.IntToBytes(protoNumber)
		content = append(content, data...)
		_, err := conn.Write(content)
		if err != nil {
			fmt.Println("send msg error ", err.Error())
		}
	} else {
		fmt.Println("kcp connect is nil ", time.Now().String())
	}
}

func handleKcpConn(conn *kcp.UDPSession) {
	buf := make([]byte, 65535)
	for {
		num, err := conn.Read(buf)
		if err != nil {
			fmt.Println("接收数据失败!", err)
			return
		}
		fmt.Printf("接收服务端数据长度：%d, 数据：%s\n", num, string(buf[4:num]))
		time.Sleep(time.Second)
	}
}
