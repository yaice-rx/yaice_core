package sort

import "github.com/satori/go.uuid"

type Town struct {
	guid 		uuid.UUID
	playerGuid 	uuid.UUID
	X 			uint16
	Y 			uint16
	Level 		uint8
}

func InitTown() *Town{
	
	return &Town{
		guid:uuid.Must(uuid.NewV4()),
	}
}
