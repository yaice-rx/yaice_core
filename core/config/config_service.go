package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

//公共配置
type commonConfigModel struct {
	mutex             sync.Mutex
	EtcdConnectString string `yaml:"EtcdConnectString"`
	EtcdNameSpace     string `yaml:"EtcdNameSpace"`
	PortStart         int    `yaml:"PortStart"`
	PortEnd           int    `yaml:"PortEnd"`
}

var CommonConfigData commonConfigModel

func InitCommonConfig() {
	CommonConfigData = commonConfigModel{}
	CommonConfigData.mutex.Lock()
	defer CommonConfigData.mutex.Unlock()
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		logrus.Println(err.Error())
		return
	}
	err = yaml.Unmarshal(yamlFile, &CommonConfigData)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
}

//服务配置
type ServiceModel struct {
	name        string //服务器名称
	groupId     string //服务器组编号
	serverExtra string
	outHost     string //内部连接ip
	outPort     int    //内部连接端口
	inHost      string //外部连接ip
	inPort      int    //内部连接端口
	IsConnect   bool   //是否需要连接
}

var confServiceData *ServiceModel

func InitServiceConf(name string, groupId string, serverExtra string, inHost string, outHost string, isConn bool) {
	confServiceData = &ServiceModel{
		name:        name,
		groupId:     groupId,
		serverExtra: serverExtra,
		inHost:      inHost,
		outHost:     outHost,
		IsConnect:   isConn,
	}
}

func GetName() string {
	return confServiceData.name
}

func GetGroupId() string {
	return confServiceData.groupId
}

func GetServerExtra() string {
	return confServiceData.serverExtra
}

func GetOutHost() string {
	return confServiceData.outHost
}

func GetOutPort() int {
	return confServiceData.outPort
}

func GetInHost() string {
	return confServiceData.inHost
}

func GetInPort() int {
	return confServiceData.inPort
}

func GetIsConn() bool {
	return confServiceData.IsConnect
}

func SetInPort(port int) {
	confServiceData.inPort = port
}

func SetOutPort(port int) {
	confServiceData.outPort = port
}

func GetServiceConfData() ServiceModel {
	if nil == confServiceData {
		return ServiceModel{}
	}
	return *confServiceData
}
