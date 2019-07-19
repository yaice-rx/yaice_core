package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

type DBModel struct {
	mutex      sync.Mutex
	db  *gorm.DB
}

var _dbModel *DBModel

func Init()*DBModel{
	if _dbModel == nil{
		//连接句柄
		dbConn,_ := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/war_kwar")
		//最大空闲连接数
		dbConn.DB().SetMaxIdleConns(10)
		//最大连接数
		dbConn.DB().SetMaxOpenConns(200)
		//超时时间
		dbConn.DB().SetConnMaxLifetime(time.Hour)
		//DBModel指针
		dbModel := new(DBModel)
		dbModel.db = dbConn
		_dbModel = dbModel
	}
	return _dbModel
}

//缓存用户数据
func (db *DBModel)CachePlayerData(playerList []*Player)[]*Player{
	db.db.Find(&playerList)
	return playerList
}