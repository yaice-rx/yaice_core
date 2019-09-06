package DataBase

type Player struct {
	Id            uint32 `gorm:"primary_key;column:id"`
	PlayerGuid    string `gorm:"column:player_guid"`
	NickName      string `gorm:"column:nick_name"`
	UserName      string `gorm:"column:user_name"`
	Password      string `gorm:"column:password"`
	Lv            uint   `gorm:"column:lv"`
	VipLv         uint   `gorm:"column:vip_lv"`
	Fight         uint32 `gorm:"column:fight"`
	LastLoginTime uint32 `gorm:"column:last_login_time"`
	CreateTime    uint32 `gorm:"column:create_time"`
}

func (Player) TableName() string {
	return "k_user"
}
