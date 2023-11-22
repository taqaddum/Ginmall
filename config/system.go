package config

type sysConf struct {
	Host  string `validate:"required"`
	Port  uint16 `validate:"required_with=Host"`
	Debug bool
	Mode  string `validate:"eq=master|eq=slave"`
}
