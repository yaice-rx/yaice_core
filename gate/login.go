package gate

import (
	"io"
	"net/http"
)

func Initialize(port string, server_id string) {
	//读取配置文件中的zookeeper地址
	//监听Http服务器
	mux := http.NewServeMux()
	mux.HandleFunc("/", login)
	http.ListenAndServe(":"+port, mux)
}

func handleMux(conn io.ReadWriteCloser) {

}

func login(w http.ResponseWriter, r *http.Request) {

}
