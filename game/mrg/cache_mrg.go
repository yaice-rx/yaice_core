package mrg

import (
	"YaIce/core/model"
)

var cachePlayerList map[string]*model.Player

func InitCacheDBData() {
	//缓存用户数据
	userData := model.Init().CachePlayerData([]*model.Player{})
	for _,value := range userData{
		cachePlayerList[value.PlayerGuid] = value
	}
}

//获取用户数据
func GetCachePlayer(playerGuid string)model.Player{
	return *cachePlayerList[playerGuid]
}