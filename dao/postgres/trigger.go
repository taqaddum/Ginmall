package postgres

import (
	"GinMall/config"
	"GinMall/model"
	"GinMall/model/enum"
	"GinMall/utils"
	"fmt"
	"github.com/go-faker/faker/v4"
	"log/slog"
)

func addDefaultAdmin() error {
	var user model.User
	if num, err := db.Count(&user); err != nil || num > 0 {
		slog.Warn("用户已存在或发生未知错误", err.Error())
		return err
	}

	user.Username = "admin"
	user.Authority = enum.Admin
	passwd := faker.Password()
	user.Password = utils.BcryptPasswd(passwd)

	if _, err := db.InsertOne(&user); err != nil {
		slog.Error("添加默认用户失败", err.Error())
		panic(err)
	} else {
		fmt.Printf("用户名:admin\n密码:%s", passwd)
	}
	return nil
}

func syncSettings() error {
	var settings = struct {
		model.Preset
		key   string
		value string
	}{}

	var action = func(arr map[string]string, exist bool) error {
		session := db.NewSession()
		defer session.Close()

		if session.Begin() == nil {
			for key, val := range config.Env {
				settings.key = key
				settings.value = val

				if !exist {
					session.Insert(&settings)
				} else {
					session.Update(&settings)
				}
			}
		}
		return session.Commit()
	}

	exist, _ := db.IsTableExist("settings")
	if !exist {
		if err := db.CreateTables(&settings); err != nil {
			slog.Error("settings表创建失败", err.Error())
			return err
		}
	}
	return action(config.Env, exist)
}
