package _map

import (
	"YaIce/core/connect"
	"YaIce/game/map/sort"
	"github.com/satori/go.uuid"
	"time"
)

//单块区域的数据
type Vision struct {
	//唯一id
	guid 		string
	//观察者
	observerGuid map[string]*MapObserver //map[用户id]地图用户信息
	//资源
	resourceList []*sort.Resource
	//怪物
	monsterList []*sort.Monster
	//对应的邻居信息
	neighborList []*Vision
	//格子编号
	col  		int
	row 		int
}

type VisionList struct {
	//视野列表
	visionGuidList	[][]Vision
}

var VisionData  *VisionList

//针对地图，切分视野
func InitVision(){
	if nil != VisionData{
		widthVisionNum ,heightVisionNum := 0,0
		//宽度
		if widthGridNum % visionRangeNum == 0 {
			widthVisionNum = widthGridNum / visionRangeNum
		} else {
			widthVisionNum = widthGridNum / visionRangeNum + 1
		}
		//高度
		if 0 == heightGridNum % visionRangeNum {
			heightVisionNum = heightGridNum / visionRangeNum
		} else {
			heightVisionNum = heightGridNum / visionRangeNum + 1
		}
		v := new(VisionList)
		//初始化格子
		v.constructVision(widthVisionNum,heightVisionNum)
		//初始化邻居信息
		v.constructNeighbor(widthVisionNum,heightVisionNum)

		VisionData = v
	}
}

//构建摄像机
func (v *VisionList)constructVision(col,row int){
	//初始化格子
	for  i := 0;i < col; i++ {
		var value []Vision
		for j := 0; j < row ; j++ {
			//初始化格子
			uuid := uuid.Must(uuid.NewV4()).String()
			vision := Vision{
				guid:uuid,
				col:i,
				row:j,
			}
			value = append(value,vision)
		}
		v.visionGuidList = append(v.visionGuidList,value)
	}
}

//构建邻居信息
func (v *VisionList)constructNeighbor(col,row int){
	for i := 0;i < col;i++{
		for  j := 0; j < row ; j++  {
			//判断区域周边是否存在
			for _i := i - 1 ;_i <= i + 1 ; _i++  {
				for _j := j - 1; _j <= j + 1 ; _j++  {
					if _i < 0 || _i >= col || _j < 0 || _j >= row {
						continue
					}
					if _i == i && _j == j  {
						continue
					}
					v.visionGuidList[i][j].neighborList =
						append(v.visionGuidList[i][j].neighborList,&v.visionGuidList[_i][_j])
				}
			}
		}
	}
}

//把对应的观察者添加到对应的区域
func (v *VisionList) AddObserver(x , y int,connect *connect.PlayerConn){
	//把对应的坐标转换成对应的视野格子坐标
	col := int(x/visionRangeNum)
	row := int(y/visionRangeNum)
	Observer := &MapObserver{
		playerGuid:connect.GetPlayerGuid(),
		col:col,
		row:row,
		joinAt:time.Now().Unix(),
	}
	MapObserverList[connect.GetPlayerGuid()] = Observer
	//加入观察者流程
	//判断此前视野中是否存在此人，存在不更新 todo
	if nil == v.visionGuidList[col][row].observerGuid[Observer.playerGuid]{
		//加入新的视野
		v.visionGuidList[col][row].observerGuid[Observer.playerGuid] =  Observer
		//给玩家推送新视野的数据
		v.PushVisionDataToPlayer(col,row,connect)
	}
}

//推送视野内的数据给玩家
func (v *VisionList)PushVisionDataToPlayer(col,row int,player *connect.PlayerConn){

	//todo 推送自己信息

	//推送周围的邻居信息
	for i := 0 ; i< len(v.visionGuidList[col][row].neighborList);i++ {

	}

	//player.WriteMsg()
}

//向在此区域中的观察者成员，广播新物品的增加
func (v *VisionList)BroadcastMapData(col,row int,data []interface{}){
	for i := 0 ;i < len(data);i++ {
		switch data[i].(type) {
		//野怪
		case sort.Monster:
			break
		//资源
		case sort.Resource:
			break
		//城镇
		case sort.Town:
			break
		}
	}
	//获取周围区域数据，进行广播
	for _,player := range v.visionGuidList[col][row].observerGuid{
		//todo
		if player == nil {

		}
	}
}
