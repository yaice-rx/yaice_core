package job

import (
	"github.com/satori/go.uuid"
	"time"
)

type JobItem struct {
	guid         string
	fn           func() //调用函数
	actionTime   int64  //执行时间
	intervalTime int64  //间隔时间
	execNum      int    //执行次数
}

type Cron struct {
	cronInterface
	running bool
	entries chan *JobItem
}

var Crontab *Cron

func Init() {
	Crontab = &Cron{
		running: true,
		entries: make(chan *JobItem, 100),
	}
	go exec()
}

type cronInterface interface {
	AddCronTask(_time int64, execNum int, handler func())
}

//加入工作列表
// t = 秒
func (this *Cron) AddCronTask(_time int64, execNum int, fn_ func()) {
	job := &JobItem{
		guid:         uuid.Must(uuid.NewV4()).String(),
		fn:           fn_,
		intervalTime: _time,
		actionTime:   time.Now().Unix(),
		execNum:      execNum,
	}
	this.entries <- job
}

func exec() {
	timer := time.NewTicker(time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			for v := range Crontab.entries {
				if nil == v {
					continue
				}
				curTime := time.Now().Unix()
				if (v.actionTime + v.intervalTime) <= curTime {
					go v.fn()
					v.actionTime = curTime
					if v.execNum != -1 {
						v.execNum--
					}
				}
				if v.execNum > 0 || v.execNum == -1 {
					Crontab.entries <- v
				}
			}
		}
	}
}
