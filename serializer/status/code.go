package status

const (
	Ok = iota + 200
	UpdatePasswdOk
	NotExistIdentify
	InvalidParams = 400
	Error         = 500
)

// 用户错误
const (
	ErrorExistUser = iota + 10001
	ErrorExistEmail
	ErrorExistPhone
	ErrorNotExistUser
	ErrorRegisterUser
	ErrorPasswdCompare
	ErrorPasswdConfirm //密码不一致
	ErrorEncryption
	ErrorNotExistProduct
	ErrorNotExistAddress
	ErrorExistFavorite
	ErrorUserNotFound
	ErrorUnLogin
	ErrorTokenExpired
	ErrorUpdateUser
)
