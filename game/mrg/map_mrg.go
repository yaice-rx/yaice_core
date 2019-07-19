package mrg

import (
	"YaIce/core/connect"
	"YaIce/game/map"
	"YaIce/protobuf"
	"github.com/golang/protobuf/proto"
)

type MapMrg struct {
	_map.MapInterface
}

//加入世界地图
func JoinMapHandler(connect *connect.PlayerConn,content []byte){
	protoData := c2game.C2GJoinMap{}
	err := proto.Unmarshal(content,&protoData)
	if nil !=  err {
		//todo send error
	}

	//_map.ConversionCoord();
	//初始化大地图用户信息
	//_map.VisionData.AddObserver();
}
