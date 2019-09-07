package auth

import (
	"YaIce/core/handler"
	"YaIce/core/job"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Initialize(port string, server_id string) {
	//初始化定时器
	job.Crontab.AddCronTask(60, -1, func() {
		logrus.Println("------------1-----------", time.Now().Unix())
	})
	job.Crontab.AddCronTask(30, -1, func() {
		logrus.Println("------------2-----------", time.Now().Unix())
	})
	//注册内部路由
	registerRouter()
	//向服务中注册自己节点数据
	handler.RegisterServiceConfigData()
	//监听Http服务器
	mux := http.NewServeMux()
	mux.HandleFunc("/", login)
	http.ListenAndServe(":"+port, mux)

}

func login(w http.ResponseWriter, r *http.Request) {
	logrus.Println(r.RequestURI)
}
