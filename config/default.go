package config

var DB = &dbConf{
	"localhost",
	5432,
	"postgres",
	"postgres",
	"postgres",
	make(map[string]string),
}

var System = &sysConf{
	"localhost",
	8080,
	true,
	"master",
}

var conf = map[string]any{
	"postgres": DB,
	"system":   System,
}

// DBVersion 与当前版本匹配的数据库版本
var dbVersion = "1.0.0-alpha"

// BackendVersion 当前后端版本号
var backendVersion = "1.0.0-alpha"

// StaticVersion 与当前版本匹配的静态资源版本
var staticVersion = "1.0.0-alpha"

// Env 环境变量
var Env = map[string]string{
	"db_version":      dbVersion,
	"backend_version": backendVersion,
	"static_version":  staticVersion,
}
