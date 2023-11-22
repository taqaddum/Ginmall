package enum

type Authority uint8

const (
	Admin Authority = iota
	Normal
	Guest
	Banned
)

func (s Authority) String() string {
	var str string
	switch s {
	case Admin:
		str = "管理员"
	case Normal:
		str = "普通用户"
	case Guest:
		str = "游客"
	case Banned:
		str = "黑名单用户"
	}
	return str
}
