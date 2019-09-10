package sort

import (
	"github.com/satori/go.uuid"
)

type Resource struct {
	//guid
	Guid uuid.UUID
	//总资源量
	TotalCapacity uint32
	//剩余的资源量
	OverCapacity uint32
	//占用的玩家，不占 == 0
	PlayerGuid uuid.UUID
	//等级
	Level uint8
	//坐标
	X uint16
	Y uint16
}

//初始化资源
func InitResource() *Resource {

	return &Resource{
		Guid: uuid.Must(uuid.NewV4()),
	}
}
