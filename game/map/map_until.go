package _map

//大地图观察者
type MapObserver struct {
	playerGuid string
	col int
	row int
	joinAt int64	//加入大地图时间
}

//大地图观察者列表
var MapObserverList map[string]*MapObserver

func AddCaCheObserver(){
	
}

//将长时间没有移动过的用户T出用户列表中
func Tick(){

}

//刷新资源
func RefreshResource(){

}

//刷新野怪
func RefreshMonster(){

}

//补充资源
func SupplementResource(){

}

//补充野怪
func SupplementMonster(){

}

func ConversionCoord(x,y int )(int,int){
	col := int(x/visionRangeNum)
	row := int(y/visionRangeNum)
	return col,row
}