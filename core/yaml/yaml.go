package yaml

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type yamlModel struct {
	mutex             sync.Mutex
	EtcdConnectString string `yaml:"EtcdConnectString"`
	EtcdNameSpace     string `yaml:"EtcdNameSpace"`
	NetworkPortStart  int    `yaml:"NetworkPortStart"`
	NetworkPortEnd    int    `yaml:"NetworkPortEnd"`
	ClusterName       string `yaml:"ClusterName"`
}

var YamlDevMrg *yamlModel

func Init() {
	if YamlDevMrg == nil {
		YamlDevMrg = &yamlModel{}
		YamlDevMrg.mutex.Lock()
		defer YamlDevMrg.mutex.Unlock()
		yamlFile, err := ioutil.ReadFile("conf.yaml")
		if err != nil {
			logrus.Println(err.Error())
			return
		}
		err = yaml.Unmarshal(yamlFile, YamlDevMrg)
		if err != nil {
			fmt.Printf("Unmarshal: %v", err)
		}
	}
}
