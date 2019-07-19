package temp

import (
	"YaIce/core/common"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"sync"
)

//配置数据
type ConfigModule struct {
	rw_mutex sync.Mutex
	allianceGiftList 	[]TemplateAliianceGift
	allianceGiftLvList 	[]TempAllianceGiftLv
	allianceRankList	[]TempAllianceRank
}
//初始化配置表数据
func InitConfigData()*ConfigModule{
	confClass := new(ConfigModule)
	confClass.loadAllianceGiftData()
	confClass.loadAllianceGiftLvData()
	confClass.loadAllianceRankData()
	return confClass
}

var mutex sync.Mutex

//读取 yaml 配置文件
func ReadConfigData() ConfigStruct {
	mutex.Lock()
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	configStruct := ConfigStruct{}
	err = yaml.Unmarshal(yamlFile, &configStruct)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	mutex.Unlock()
	return configStruct
}

//加载礼物数据
func (c *ConfigModule) loadAllianceGiftData(){
	data := common.ReadCSVData("AllianceGift")
	for  i := 0; i< len(data);i++{
		giftId,_ := strconv.Atoi(data[i][0])
		giftType,_ := strconv.Atoi(data[i][2])
		chestExp,_ := strconv.Atoi(data[i][3])
		itemChestItem,_ := strconv.Atoi(data[i][4])
		tempData := TemplateAliianceGift{
			GiftId: giftId,
			GiftType :giftType,
			ChestExp :chestExp,
			ItemChestId :itemChestItem,
		}
		c.allianceGiftList = append(c.allianceGiftList,tempData)
	}
}

//联盟大礼包
func (c *ConfigModule) loadAllianceGiftLvData() {
	data := common.ReadCSVData("AllianceGiftLv")
	for  i := 0; i< len(data);i++{
		id,_ := strconv.Atoi(data[i][0])
		exp,_ := strconv.Atoi(data[i][2])
		tempData := TempAllianceGiftLv{
			ID: id,
			BigGiftPro :data[i][1],
			Exp :exp,
		}
		c.allianceGiftLvList = append(c.allianceGiftLvList,tempData)
	}
}

//联盟成员等级
func (c *ConfigModule) loadAllianceRankData() {
	data := common.ReadCSVData("AllianceRank")
	for  i := 0; i< len(data);i++{
		id,_ := strconv.Atoi(data[i][0])
		number,_ := strconv.Atoi(data[i][1])
		tempData := TempAllianceRank{
			ID: id,
			Num :number,
		}
		c.allianceRankList = append(c.allianceRankList,tempData)
	}
}