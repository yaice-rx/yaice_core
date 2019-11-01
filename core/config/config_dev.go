package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

//公共配置
type ConfigDevModel struct {
	mutex             		sync.Mutex
	EtcdConnectString 		string `yaml:"EtcdConnectString"`
	EtcdNameSpace     		string `yaml:"EtcdNameSpace"`
	NetworkPortStart        int    `yaml:"NetworkPortStart"`
	NetworkPortEnd          int    `yaml:"NetworkPortEnd"`
	ClusterName				string	`yaml:"ClusterName"`
}

var ConfDevMrg *ConfigDevModel

func InitImplEtcd() {
	if ConfDevMrg == nil{
		ConfDevMrg = &ConfigDevModel{}
		ConfDevMrg.mutex.Lock()
		defer ConfDevMrg.mutex.Unlock()
		yamlFile, err := ioutil.ReadFile("conf.yaml")
		if err != nil {
			logrus.Println(err.Error())
			return
		}
		err = yaml.Unmarshal(yamlFile, ConfDevMrg)
		if err != nil {
			fmt.Printf("Unmarshal: %v", err)
		}
	}
}

