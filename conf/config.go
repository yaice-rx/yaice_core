package conf

import (
	"sync"
)

//配置数据
type ConfigModule struct {
	mutex          sync.Mutex
	rw_mutex       sync.Mutex
	TempGiftList   []AliianceGiftModel
	TempGiftLvList []AllianceGiftLvModel
	TempRankList   []AllianceRankModel
}

var ConfigEntity *ConfigModule

//初始化配置表数据
func InitCSVConfigData() {
	ConfigEntity = new(ConfigModule)
	ConfigEntity.loadAllianceGiftData()
}

//加载礼物数据
func (c *ConfigModule) loadAllianceGiftData() {
	/*data := common.ReadCSVData("AllianceGift")
	for  i := 0; i< len(data);i++{
		giftId,_ := strconv.Atoi(data[i][0])
		giftType,_ := strconv.Atoi(data[i][2])
		chestExp,_ := strconv.Atoi(data[i][3])
		itemChestItem,_ := strconv.Atoi(data[i][4])
		tempData := AliianceGiftModel{
			GiftId: giftId,
			GiftType :giftType,
			ChestExp :chestExp,
			ItemChestId :itemChestItem,
		}
		c.TempGiftList = append(c.TempGiftList,tempData)
	}*/
}

//联盟大礼包
func (c *ConfigModule) loadAllianceGiftLvData() {
}

//联盟成员等级
func (c *ConfigModule) loadAllianceRankData() {
}
