package postgres

import (
	"GinMall/config"
	_ "github.com/lib/pq"
	"log/slog"
	"sync"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var db *xorm.Engine
var once sync.Once

func InitDB() *xorm.Engine {
	once.Do(func() {
		var err error
		db, err = initDatabase()
		if err != nil {
			slog.Error("数据库初始化失败", err.Error())
		}
	})
	return db
}

func initDatabase() (*xorm.Engine, error) {
	database, err := xorm.NewEngine("postgres", config.DB.DSN())
	if err != nil {
		return nil, err
	}

	database.SetMapper(names.GonicMapper{})
	database.SetMaxIdleConns(10)
	database.SetMaxOpenConns(50)
	database.SetConnMaxLifetime(4 * time.Hour)
	database.ShowSQL(config.System.Debug)
	database.Logger().SetLevel(log.LOG_DEBUG)

	if err := database.Ping(); err != nil {
		slog.Error("数据库连接失败", err.Error())
		return nil, err
	}

	migration()

	if err := addDefaultAdmin(); err != nil {
		slog.Warn("添加默认用户失败", err.Error())
	}
	return database, nil
}
