package main

import (
	"GinMall/config"
	"fmt"
	"log/slog"
)

func main() {
	if err := config.InitConf(); err != nil {
		slog.Debug("配置初始化失败", err.Error())
	}
	fmt.Println(config.DB.DSN())
	fmt.Scanln()
}
