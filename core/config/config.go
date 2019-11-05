package config

import (
	"os"
	"strconv"
	"strings"
)

type ModuleConfig struct {
	Pid                string   //服务进程编号
	TypeName           string   //服务类型
	GroupName          string   //服务组编号
	ConnServerNameList []string //限制连接的服务
	OutHost            string   //外部连接ip
	OutPort            string   //外部连接端口
	InHost             string   //内部连接ip
	InPort             string   //内部连接端口
	ConnectPath        string   //连接路径
	HttpPort           string
}

var Config *ModuleConfig

func Init(_type string, groupName string, connStr string, inHost string, outHost string, port string) {
	Config = &ModuleConfig{
		Pid:                strconv.Itoa(os.Getpid()),
		TypeName:           _type,
		GroupName:          groupName,
		ConnServerNameList: strings.Split(connStr, ","),
		InHost:             inHost,
		OutHost:            outHost,
		HttpPort:           port,
	}
}
