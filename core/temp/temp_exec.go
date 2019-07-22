package temp

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)

//读取 yaml 配置文件
func (c *ConfigModule)loadYamlConfigData()  {
	c.mutex.Lock()
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		logrus.Println(err.Error())
		return
	}
	configData := yamlConfigData{}
	err = yaml.Unmarshal(yamlFile, &configData)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	c.mutex.Unlock()
	c.YamlConfigData = configData
}

//加载礼物数据
func (c *ConfigModule) loadAllianceGiftData(){
	data := ReadCSVData("AllianceGift")
	for  i := 0; i< len(data);i++{
		giftId,_ := strconv.Atoi(data[i][0])
		giftType,_ := strconv.Atoi(data[i][2])
		chestExp,_ := strconv.Atoi(data[i][3])
		itemChestItem,_ := strconv.Atoi(data[i][4])
		tempData := templateAliianceGift{
			GiftId: giftId,
			GiftType :giftType,
			ChestExp :chestExp,
			ItemChestId :itemChestItem,
		}
		c.TempGiftList = append(c.TempGiftList,tempData)
	}
}

//联盟大礼包
func (c *ConfigModule) loadAllianceGiftLvData() {
	data := ReadCSVData("AllianceGiftLv")
	for  i := 0; i< len(data);i++{
		id,_ := strconv.Atoi(data[i][0])
		exp,_ := strconv.Atoi(data[i][2])
		tempData := tempAllianceGiftLv{
			ID: id,
			BigGiftPro :data[i][1],
			Exp :exp,
		}
		c.TempGiftLvList = append(c.TempGiftLvList,tempData)
	}
}

//联盟成员等级
func (c *ConfigModule) loadAllianceRankData() {
	data := ReadCSVData("AllianceRank")
	for  i := 0; i< len(data);i++{
		id,_ := strconv.Atoi(data[i][0])
		number,_ := strconv.Atoi(data[i][1])
		tempData := tempAllianceRank{
			ID: id,
			Num :number,
		}
		c.TempRankList = append(c.TempRankList,tempData)
	}
}

