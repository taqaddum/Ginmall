package service

import (
	"GinMall/dao"
	"GinMall/model"
	"GinMall/model/enum"
	"GinMall/serializer/request"
	"GinMall/serializer/response"
	"GinMall/serializer/status"
	"GinMall/utils"
)

type UserSrvApi interface {
	RegisterUser(req *request.RegisterRequest) *response.BasicResponse
	LoginUser(req *request.LoginRequest) *response.BasicResponse
	GetUserDetail(uname string) *response.BasicResponse
	UpdateUserInfo(req *request.UpdateRequest) *response.BasicResponse
}

type UserService struct {
	DAO dao.UserDBApi
}

func (srv UserService) RegisterUser(req *request.RegisterRequest) *response.BasicResponse {
	if srv.DAO.IsDuplicate(req.Username) {
		return response.FailedCode(status.ErrorExistUser)
	}

	if req.Password != req.ConfirmPwd {
		return response.FailedCode(status.ErrorPasswdConfirm)
	}

	err := srv.DAO.Add(&model.User{Username: req.Username, Password: req.Password, Phone: req.Phone})
	if err != nil {
		return response.FailedCode(status.ErrorRegisterUser)
	}
	return response.SuccessMsg("注册成功")
}

func (srv UserService) LoginUser(req *request.LoginRequest) *response.BasicResponse {
	user, err := srv.DAO.GetByUsername(req.Username)
	if err != nil {
		return response.FailedCode(status.ErrorUserNotFound)
	}

	if utils.CheckPasswd(user.Password, req.Password) == false {
		return response.FailedCode(status.ErrorPasswdCompare)
	}
	token := utils.SignedToken(user.ID, user.Username, user.Authority)
	return response.SuccessData(token)
}

func (srv UserService) GetUserDetail(uname string) *response.BasicResponse {
	user, err := srv.DAO.GetByUsername(uname)
	if err != nil {
		return response.FailedCode(status.ErrorUserNotFound)
	}
	return response.SuccessData(user)
}

func (srv UserService) UpdateUserInfo(req *request.UpdateRequest) *response.BasicResponse {
	if utils.Verify(req.Password) != nil {
		req.Password = ""
	}

	err := srv.DAO.Update(&model.User{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Gender:   enum.GetGender(req.Gender),
	})

	if err != nil {
		return response.FailedCode(status.ErrorUpdateUser)
	}
	return response.Success()
}
