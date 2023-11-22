package config

import (
	"GinMall/utils"
	"flag"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
)

func InitConf() error {
	var dir string
	flag.StringVar(&dir, "conf", ".", "指定配置文件所在目录")
	flag.Parse()

	path := filepath.Join(filepath.Dir(dir), "configure.toml")
	if err := generateDefaultConf(path); err != nil {
		slog.Debug("生成默认配置失败", err.Error())
		return err
	}

	if err := readAndValidateConf(path); err != nil {
		slog.Debug("读取配置失败", err.Error())
		return err
	}
	return nil
}

func generateDefaultConf(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		slog.Debug("未检测到配置文件，生成默认配置...")
		tomlByte, err := toml.Marshal(conf)
		if err != nil {
			slog.Error("序列化配置项失败！", err.Error())
			return err
		}
		utils.WriteConf(path, tomlByte)
	}
	return nil
}

func readAndValidateConf(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("配置文件解析失败", err.Error())
		return err
	}

	for key, value := range conf {
		if viper.IsSet(key) {
			viper.UnmarshalKey(key, value)
			if err := utils.Verify(value); err != nil {
				slog.Error(key, "存在不合法数据", err.Error())
				return err
			}
		} else {
			slog.Warn(key, "%s配置项不存在", key)
		}
	}
	return nil
}
