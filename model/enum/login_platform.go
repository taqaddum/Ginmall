package enum

type Platform int

const (
	Local Platform = iota
	Email
	Phone
	QQ
	WeChat
)

func (p Platform) String() string {
	var str string
	switch p {
	case Local:
		str = "用户登录"
	case Email:
		str = "邮箱登录"
	case Phone:
		str = "手机登录"
	case QQ:
		str = "QQ登录"
	case WeChat:
		str = "微信登录"
	}
	return str
}
