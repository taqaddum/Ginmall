package request

type UpdateRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Avatar   string `form:"avatar"`
	Gender   string `form:"gender"`
}
