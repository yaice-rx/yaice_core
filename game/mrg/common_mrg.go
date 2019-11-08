package mrg

import (
	"YaIce/core/model"
	"YaIce/protobuf/outer_proto"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

//处理ping包
func PingHandler(connect *model.Conn, content []byte) {

}

func LoginHandler(connect *model.Conn, content []byte) {
	data := outer_proto.C2GLogin{}
	err := proto.Unmarshal(content, &data)
	if err != nil {
		logrus.Println("Unmarshal data error: ", err)
	}
	//连接
	//cluster.ClusterMrg.ConnList[yaml.YamlDevMrg.ClusterName]["center"]["auth"]
}
