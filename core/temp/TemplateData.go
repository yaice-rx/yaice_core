package temp
//配置文件数据
type ConfigStruct struct {
	ZooConnectString string `yaml:"ZooConnectString"`
	ZooNameSpace string	`yaml:"ZooNameSpace"`
	TcpServicePortStart int	`yaml:"TcpServicePortStart"`
	TcpServicePortEnd	int	`yaml:"TcpServicePortEnd"`
	ServerPingServerInterval	int	`yaml:"ServerPingServerInterval"`
	ExcelPath	string	`yaml:"ExcelPath"`
	PlayerCacheUpdateInterval	int	`yaml:"PlayerCacheUpdateInterval"`
	PingMapCameraTimeout 	int	`yaml:"PingMapCameraTimeout"`
	OneLittleGridLength	int	`yaml:"OneLittleGridLength"`
	OneVisionLittleGridNum 	int `yaml:"OneVisionLittleGridNum"`
	WidthGridNum	int `yaml:"WidthGridNum"`
	HeightGridNum	int `yaml:"HeightGridNum"`
}

//联盟礼物
type TemplateAliianceGift struct {
	GiftId int
	GiftType int
	ChestExp int
	ItemChestId int
}

//联盟
type TempAllianceGiftLv struct {
	ID int
	BigGiftPro string
	Exp int
}

//联盟成员等级
type TempAllianceRank struct {
	ID int
	Num int
}

//地图信息
type TempArea struct {
	ID 				int
	UnlockBuilding 	string
	AmountMax 		string
	AreaRange 		string
	Troops			int
	Hero			int
	INT_FoodReq 	int
	INT_StoneReq 	int
	INT_WoodReq 	int
	INT_IronReq		int
	INT_CoinReq 	int
}

//军械
type TempArmaments struct {
	ID	int
	Type	int
	Level	int
	Atk		int
	Def		int
	Hp		int
	Skill	int
	Unlock	string
	PreBuild	int //升级需求建筑条件
	FoodReq		int //消耗粮食
	StoneReq	int //消耗石料
	WoodReq		int	//消耗木材
	IronReq		int	//消耗铁矿
	CoinReq		int //消耗金币
	Item1Req	int //需求道具
	BuildTime	int //升级时间（秒）
	Power		int //战斗力提升
}

//军械库等级
type TempArmory struct {
	ID	int
	LevelUp		string
	AddBuff		string
}

//箭塔
type TempArrowTower struct {
	ID	int
	Atk	int	//攻击
	Def	int	//防御
	Hp	int	//生命
	PreBuild	int	//升级需求城墙等级
	FoodReq		int	//消耗粮食
	StoneReq	int	//消耗石料
	WoodReq		int	//消耗木材
	IronReq		int	//消耗铁矿
	CoinReq		int	//消耗金币
	Item1Req	int	//需求道具
	BuildTime	int	//升级时间（秒）
	Power		int	//战斗力提升
}

//集会所配置
type TempAssembly struct {
	ID	int //##集会所配置
	LevelUp	int //集结部队上限
	ExtraBuff	int //满级效果（BUFFID#效果值万分比）
}

//兵营配置
type TempBarracks struct {
	ID	int //##level
	TrainingCnt int //训练数量
	AddBuff		int	//附加buff（ID#值，万分比）
}

//Buffer
type TempBuffDes struct {
	ID		int	//##BUFFID
	ID1		int	//字段ID
	Type	int	//数值类型（0-特殊，1-数值，2-百分比）
}

//主城
type TempCastle struct {
	ID			int //##市政厅相关功能
	MarchCnt	int	//行军队列
	HelpCnt		int	//帮助次数
	TroopNum	int	//军团规模
	HeroNum		int	//军团中英雄数量
	FoodServ	int	//升级奖励粮食
	StoneServ	int	//升级奖励石材
	WoodServ	int	//升级奖励木
	IronServ	int	//升级奖励铁矿
	CoinServ	int	//升级奖励金币
	Item		string	//升级奖励道具
	AllianceGift	int		//升级提供联盟礼物
}

//建筑
type TempCityBuilding struct {
	ID	int	//编号
	UID	int	//建筑唯一配置ID
	Type	int	//建筑类型（1政府建筑2军事建筑3资源建筑4守护神建筑）
	ID1	int	//建筑等级
	STR_DefaultPos	int //建筑大小（长#宽）
	ReBuild	int	//能否重复修建（0不能1能）
	BuildNumMax	int	//修建数量上限
	INT_CanRemove	int	//能否拆除（0不能1能）
	INT_RemoveCost	int	//拆除消耗道具ID
	INT_PreBuild1Type	int //前置建筑（ID#等级；…)
	INT_FoodReq	int	//粮食消耗
	INT_WoodReq	int	//木材消耗
	INT_StoneReq	int	//石头消耗
	INT_IronReq	int	//铁矿消耗
	INT_CoinReq	int	//金币消耗
	INT_Item1Req	int	//需求道具（ID#数量；…）
	INT_BuildTime	int	//修建时间（秒）
	INT_ExpAward	int	//经验
	INT_PowerAward	int	//战力
}

//使馆
type TempEmbassy struct {
	ID	int	//##使馆配置
	LevelUp	int	//援军上限
	ExtraBuff	string //满级效果（BUFFID#效果值万分比）
}

//装备
type TempEquipment struct {
	ID	int
	EquipID	int
	Type	int
	Quality	int
	LevelLimit	int
	Buff	string
	PreEquip	int
	Upgrade		string
	Forging_Time	int
	Coin_Cost		int
	Equip_Series	int
	Not_forgable	int
	Exp		int
}

//装备套装
type TempEquipSource struct {
	ID	int
	Source	string
}

//联盟权限
type TempFactionChmod struct {
	ID	int //帮会阶级ID
	GetGift	int	//收取礼物
	MailEveryOne	int	//全盟发信
	UpMemberStep	int	//提升成员阶级
	DownMemberStep	int	//降低成员阶级
	OutMember		int	//驱逐成员
	InviteJoin		int	//邀请入盟
	HandleJoinAsk	int	//管理联盟申请
	EditAnnouncement	int	//编辑盟内公告
	EditForeignAnnouncement	int	//编辑对外公告
	EditDeclaration	int	//编辑联盟宣言
	DelForeignMsg	int	//删除外交留言
	EditBookMark	int	//编辑联盟书签
	EditBadge		int	//变更徽章
	EditNameCode	int	//变更联盟名字代号
	SelectLanguage	int	//选择主要语言
	OpenRecruit		int	//公开招募
	Abdicate		int	//盟主让位
	RemoveFaction	int	//解散联盟
	ChangeServerId	int	//变更所属王国
	RecallChairmanTime	int	//发起罢免盟主时间（天）
}

//锻造厂
type TempFactory struct {
	ID	int
	ForgingSpeed int	//装备锻造速度（buffid201）增加（万分比）
}

//农场
type TempFarm struct {
	ID	int	//##农场
	Produce	int	//粮食产量（小时）
	Capacity	int //粮食容量
}

//攻击阵型
type TempFormationTarget struct {
	ID	int	//id
	Formation1	int //阵型1(1步兵阵2骑兵阵3弓兵阵）
	AtkProportion1 string	//兵种被攻击顺序(默认顺序：步骑弓）
}

//常量
type TempGdConst struct {
	ID	int	//##描述
	INT_value	int//数值
	StrValue	string//字符串
}

//宝石属性
type TempGemProperty struct {
	ID	int
	Property string //宝石属性（buffid#万分比值;)
}

//仓库
type TempGranary struct {
	ID	int	//##仓库内容
	INT_NormalCnt	int//仓库容量（除金币）
	INT_CoinCnt int//金币容量
}

//英雄装备
type TempHeroEquip struct {
	ID	int
	Quality	int
	Level	int
	Dungeon	string
	Formula	string
	Scrap	string
	Str		int
	Agi		int
	Int		int
	Atk		int
	Pdef	int
	Mag		int
	Mdef	int
	Hp		int
	Ppenetrate	int
	Mpenetrate	int
	Crit		int
	HpRestore	int
	InitialAnger	int
	AngerRecovery	int
}

//医院
type TempHospital struct {

}

//初始数据
type TempInitial struct {

}

//部件
type TempItem struct {

}

//宝箱
type TempItemChest struct {

}

//商店
type TempItemShop struct {

}

//研究院属性
type TempLibrary struct {

}

//领主
type TempLordLevel struct {

}

//邮件
type TempMail	struct {

}

//庄园
type TempManor	struct {

}

//地图资源等级刷新范围表
type TempMapLv	struct {

}

//地图资源刷新
type TempMapTile struct {

}

//市集
type TempMarket struct {

}

//矿场
type TempMine struct {

}

//采石场
type TempQuarry struct {

}

//任务
type TempQuest struct {

}

//怪物刷新配置
type TempRallyMonster struct {

}

//学院子类
type TempResearch struct {

}

//学院大类
type TempResearchClass struct {

}

//钻石消耗资源
type TempResexchg struct {

}

//伐木场
type TempSawmill struct {

}

//技能
type TempSkill struct {

}

//套装
type TempSpecialEquip struct {

}

//天赋
type TempTalent struct {

}

//天赋类型
type TempTalentClass struct {

}

//哨塔
type TempTower struct {

}

//城墙
type TempWalls struct {

}

//主城等级
type TempWarFever struct {

}

//士兵
type TempWarriorData struct {

}