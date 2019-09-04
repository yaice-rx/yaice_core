package mrg

import (
	"YaIce/core/model"
)

var cachePlayerList map[string]*model.Player

func InitCacheDBData() {
}

//获取用户数据
func GetCachePlayer(playerGuid string) model.Player {
	return *cachePlayerList[playerGuid]
}
