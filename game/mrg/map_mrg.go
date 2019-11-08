package mrg

import (
	"YaIce/core/model"
	"YaIce/game/map"
)

type MapMrg struct {
	_map.MapInterface
}

//加入世界地图
func JoinMapHandler(connect *model.Conn, content []byte) {
	//_map.ConversionCoord();
	//初始化大地图用户信息
	//_map.VisionData.AddObserver();
}
