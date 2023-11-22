package dao

import (
	"GinMall/model"
	"GinMall/utils"
	"errors"
	"log/slog"
	"xorm.io/xorm"
)

type UserDBApi interface {
	IsDuplicate(name string) bool
	Add(user *model.User) error
	Update(user *model.User) error
	Remove(user *model.User) error
	GetAll(user model.User) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
}

type UserDAO struct {
	DB *xorm.Engine
}

func (dao UserDAO) IsDuplicate(name string) bool {
	has, err := dao.DB.Where("username=?", name).Exist()
	if err != nil {
		slog.Debug("查询失败", err.Error())
	}
	return has
}

func (dao UserDAO) Add(user *model.User) error {
	has, _ := dao.DB.Where("username=?", user.Username).Exist()
	if has {
		return errors.New("用户已存在")
	}

	affect, err := dao.DB.InsertOne(user)
	if err != nil {
		slog.Debug("记录插入失败", err.Error())
	} else {
		slog.Debug("%d条记录已插入", affect)
	}
	return err
}

// Update
//
//	@Description: 更新用户信息
//	@receiver dao 用户数据访问层实例
//	@param user 用户信息
//	@return error
func (dao UserDAO) Update(user *model.User) error {
	affect, err := dao.DB.Where("username=?", user.Username).Omit("id", "username", "authority").Update(user)
	if affect > 0 && err == nil {
		slog.Debug("%d条记录已更新", affect)
	} else {
		slog.Debug("记录未更新", err.Error())
	}
	return err
}

func (dao UserDAO) Remove(user *model.User) error {
	affect, err := dao.DB.Delete(user)
	if affect > 0 && err == nil {
		slog.Debug("%d条记录已删除", affect)
	} else {
		slog.Debug("记录未删除", err.Error())
	}
	return err
}

func (dao UserDAO) GetByUsername(username string) (*model.User, error) {
	user := &model.User{Username: username}
	exist, err := dao.DB.Get(user)
	if exist {
		return user, nil
	}
	return nil, err
}

func (dao UserDAO) GetAll(user model.User) (*model.User, error) {
	exist, err := dao.DB.Get(&user)
	if exist {
		return utils.CopyStruct(&user).(*model.User), err
	}
	return nil, err
}
