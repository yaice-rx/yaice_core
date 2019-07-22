package _map

import (
	"YaIce/core/temp"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

//地形
var MapTerrainData [][]int

//初始化类型
func InitTerrain(){
	txtData := temp.ReadTXTData("WorldMapPosition")
	if nil == txtData {
		logrus.Error("WorldMapPosition data nil")
		panic("WorldMapPosition data nil")
	}
	for i := 0 ; i < len(txtData);i++{
		str := strings.Split(txtData[i],",")
		var terrainX []int
		for j := 0;j < len(str);j++ {
			val,_ := strconv.Atoi(str[j])
			terrainX = append(terrainX,val)
		}
		MapTerrainData = append(MapTerrainData,terrainX)
	}
}

//地形检测
func CheckPostion(x,y int)bool{
	for i := x -1; i <= x + 1; i++{
		for y := y-1; y <= y + 1; y++ {
			if x < 0 || x >= 1000 || y < 0 || y >= 1000 {
				break
			}
			if 1 == MapTerrainData[i][y]{
				return false
			}
		}
	}
	return true
}