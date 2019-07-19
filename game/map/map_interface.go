package _map

type MapInterface interface {
	//初始化大地图数据
	Init()
	//开始行军
	StartMarch()
	//到达目标点
	ArrivalTarget()
	//回家
	ComeBackHome()
}

const (
	//攻击主城
	Map_AttackTown = iota
	//采集
	Map_Collect
	//攻打怪物
	Map_AttackMonster
)

//地图状态机
type MapStateMachine interface {
	//开始
	Start()
	//执行
	Exec()
	//退出
	Exit()
}

