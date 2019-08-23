package config

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"sync"
)

//配置文件数据
type YamlModel struct {
	EtcdConnectString         string `yaml:"EtcdConnectString"`
	EtcdNameSpace             string `yaml:"EtcdNameSpace"`
	PortStart                 int    `yaml:"PortStart"`
	PortEnd                   int    `yaml:"PortEnd"`
	ServerPingServerInterval  int    `yaml:"ServerPingServerInterval"`
	ExcelPath                 string `yaml:"ExcelPath"`
	PlayerCacheUpdateInterval int    `yaml:"PlayerCacheUpdateInterval"`
	PingMapCameraTimeout      int    `yaml:"PingMapCameraTimeout"`
	OneLittleGridLength       int    `yaml:"OneLittleGridLength"`
	OneVisionLittleGridNum    int    `yaml:"OneVisionLittleGridNum"`
	WidthGridNum              int    `yaml:"WidthGridNum"`
	HeightGridNum             int    `yaml:"HeightGridNum"`
}

type yamlEntity struct {
	mutex sync.Mutex
	Data  YamlModel
}

//yaml配置文件系统
var yamlConfSystem *yamlEntity

func YamlInit() {
	yamlConfSystem = &yamlEntity{}
	yamlConfSystem.initData()
}

func GetYamlData() YamlModel {
	return yamlConfSystem.Data
}

//读取 yaml 配置文件
func (c *yamlEntity) initData() {
	c.mutex.Lock()
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		logrus.Println(err.Error())
		return
	}

	configData := YamlModel{}
	err = yaml.Unmarshal(yamlFile, &configData)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	c.mutex.Unlock()
	c.Data = configData
}
