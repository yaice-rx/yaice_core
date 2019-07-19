package temp

import (
	"sync"
)

//配置数据
type ConfigModule struct {
	mutex sync.Mutex
	rw_mutex sync.Mutex
	YamlConfigData	yamlConfigData
	TempGiftList 	[]templateAliianceGift
	TempGiftLvList 	[]tempAllianceGiftLv
	TempRankList	[]tempAllianceRank
}

var ConfigCacheData *ConfigModule

//初始化配置表数据
func InitConfigData(){
	confClass := new(ConfigModule)
	confClass.loadYamlConfigData()
	confClass.loadAllianceGiftData()
	confClass.loadAllianceGiftLvData()
	confClass.loadAllianceRankData()
	ConfigCacheData = confClass
}
