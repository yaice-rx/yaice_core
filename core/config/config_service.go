package config

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
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

var CommonConfigHandler commonConfigModel

func InitCommonConfig() {
	CommonConfigHandler = commonConfigModel{}
	CommonConfigHandler.mutex.Lock()
	defer CommonConfigHandler.mutex.Unlock()
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		logrus.Println(err.Error())
		return
	}
	err = yaml.Unmarshal(yamlFile, &CommonConfigHandler)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
}

//服务配置
type ServiceConfigModel struct {
	Pid      string //服务器单独进程编号
	Name     string //服务器名称
	GroupId  string //服务器组编号
	ConnList []string
	OutHost  string //内部连接ip
	outPort  int    //内部连接端口
	InHost   string //外部连接ip
	InPort   int    //内部连接端口
}

var ConfServiceHandler *ServiceConfigModel

func InitServiceConf(name string, groupId string, connStr string, inHost string, outHost string) {
	pid, _ := uuid.NewV4()
	connList := strings.Split(connStr, ",")
	ConfServiceHandler = &ServiceConfigModel{
		Pid:      pid.String(),
		Name:     name,
		GroupId:  groupId,
		ConnList: connList,
		InHost:   inHost,
		OutHost:  outHost,
	}
}
func (this *ServiceConfigModel) GetPid() string {
	return this.Pid
}

func (this *ServiceConfigModel) GetName() string {
	return this.Name
}

func (this *ServiceConfigModel) GetGroupId() string {
	return this.GroupId
}

func (this *ServiceConfigModel) GetConnList() []string {
	return this.ConnList
}

func (this *ServiceConfigModel) GetOutHost() string {
	return this.OutHost
}

func (this *ServiceConfigModel) GetOutPort() int {
	return this.outPort
}

func (this *ServiceConfigModel) GetInHost() string {
	return this.InHost
}

func (this *ServiceConfigModel) GetInPort() int {
	return this.InPort
}

func (this *ServiceConfigModel) SetInPort(port int) {
	this.InPort = port
}

func (this *ServiceConfigModel) SetOutPort(port int) {
	this.outPort = port
}

func (this *ServiceConfigModel) GetServiceConfData() ServiceConfigModel {
	if nil == ConfServiceHandler {
		return ServiceConfigModel{}
	}
	return *ConfServiceHandler
}
