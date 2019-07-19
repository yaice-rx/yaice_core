package sort

import "github.com/satori/go.uuid"

type Monster struct {
	//guid
	Guid 			string
	//总血量值
	TotalBooldVol  	uint32
	//剩余血量值
	OverBooldVol	uint32
	//占用的玩家，不占 == 0
	PlayerGuid 		uuid.UUID
	//等级
	Level 			uint8
	//坐标
	X 				uint16
	Y 				uint16
}

var  MonsterList   map[uuid.UUID]Monster

func InitMonster()*Monster{

	return &Monster{
		Guid:uuid.Must(uuid.NewV4()).String(),
	}
}



