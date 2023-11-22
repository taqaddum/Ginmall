package postgres

import (
	"GinMall/model"
	"log/slog"
)

func migration() {
	if err := syncSettings(); err != nil {
		slog.Error("环境变量同步失败！", err.Error())
		panic(err)
	}

	err := db.Sync2(&model.User{}, &model.Banner{}, &model.Address{}, &model.CartItem{}, &model.Order{})
	if err != nil {
		slog.Error("数据模型同步失败！", err.Error())
		panic(err)
	}
}
