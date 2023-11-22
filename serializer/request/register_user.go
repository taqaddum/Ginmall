package request

type RegisterRequest struct {
	Phone      string `form:"phone" binding:"required"`
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password" binding:"required"`
	ConfirmPwd string `form:"confirm_password" binding:"required"`
}
