package status

var Msg = map[int]string{
	Ok:                   "操作成功",
	Error:                "操作失败",
	UpdatePasswdOk:       "密码修改成功",
	NotExistIdentify:     "标识不存在",
	InvalidParams:        "无效参数",
	ErrorExistEmail:      "邮箱已存在",
	ErrorExistPhone:      "手机号已存在",
	ErrorNotExistUser:    "用户不存在",
	ErrorRegisterUser:    "注册失败",
	ErrorExistUser:       "用户已存在",
	ErrorPasswdCompare:   "密码错误",
	ErrorPasswdConfirm:   "两次密码不一致",
	ErrorEncryption:      "加密失败",
	ErrorNotExistProduct: "商品不存在",
	ErrorNotExistAddress: "地址不存在",
	ErrorExistFavorite:   "已收藏",
	ErrorUserNotFound:    "用户查询失败",
	ErrorUnLogin:         "未登录或非法访问",
	ErrorTokenExpired:    "授权已过期",
	ErrorUpdateUser:      "用户信息更新失败",
}
