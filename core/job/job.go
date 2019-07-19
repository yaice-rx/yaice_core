package job

import (
	"github.com/satori/go.uuid"
	"time"
)

type JobItem struct{
	guid 	string
	fn		func()	//调用函数
	actionTime  int64	//执行时间
	intervalTime int64	//间隔时间
}

var JobList []*JobItem

//加入工作列表
// t = 秒
func JoinJob(t int64,fn_ func()){
	job := &JobItem{
		guid:uuid.Must(uuid.NewV4()).String(),
		fn:fn_,
		intervalTime:t,
		actionTime:time.Now().Unix(),
	}
	JobList = append(JobList, job)
}

func CallJob(){
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		<- t.C
		for _,v := range JobList{
			if nil == v {
				continue
			}
			curTime := time.Now().Unix()
			if (v.actionTime + v.intervalTime) >= curTime{
				go v.fn()
				v.actionTime = curTime
			}
		}
	}
}