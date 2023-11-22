package tests

import (
	"GinMall/config"
	"log"
	"testing"
)

func TestInitConf(t *testing.T) {

	if err := config.InitConf(); err != nil {
		log.Println("配置初始化失败", err.Error())
	} else {
		log.Println("数据库地址：", config.DB.DSN())
	}
}
