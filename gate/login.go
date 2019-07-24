package gate

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go"
	"io"
	"net/http"
)

func Initialize(port string ,server_id string){
	//读取配置文件中的zookeeper地址
	//监听Http服务器
	mux := http.NewServeMux()
	mux.HandleFunc("/",  login)
	http.ListenAndServe(":"+port,mux)
}

func handleMux(conn io.ReadWriteCloser) {

}

func login(w http.ResponseWriter,r *http.Request){

}
//初始化监听内网端口
func InitInteralPort(port string){
	lis, err := kcp.ListenWithOptions(port, nil, 10, 3)
	if err != nil {
		logrus.Error("")
		return
	}
	for {
		if conn, err := lis.AcceptKCP(); err == nil {
			go handleMux(conn)
		} else {
			fmt.Println(err.Error())
		}
	}
}


