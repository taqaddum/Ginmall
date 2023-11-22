package dao

import (
	"GinMall/model"
	"GinMall/utils"
	"xorm.io/xorm"
)

type BannerDBApi interface {
	List(limit int, offset int) ([]*model.Banner, error)
	Total() (int, error)
	Get(banner model.Banner) (*model.Banner, error)
}

type BannerDAO struct {
	DB *xorm.Engine
}

func (dao BannerDAO) Get(banner model.Banner) (*model.Banner, error) {
	exists, err := dao.DB.Get(&banner)
	if exists {
		return utils.CopyStruct(&banner).(*model.Banner), err
	}
	return nil, err
}

func (dao BannerDAO) List(limit int, offset int) ([]*model.Banner, error) {
	var banners = make([]*model.Banner, 0)
	err := dao.DB.Limit(limit, offset).Find(&banners)
	return banners, err
}

func (dao BannerDAO) Total() (int, error) {
	var banner model.Banner
	num, err := dao.DB.Count(&banner)
	return int(num), err
}
