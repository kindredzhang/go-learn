package manage

import (
	"errors"
	"go-learn/global"
	"go-learn/model/common"
	"go-learn/model/manage"
	manageReq "go-learn/model/manage/request"
	"go-learn/utils"
	"strconv"
	"time"
)

type ManageGoodsCategoryService struct {
}

func (s ManageGoodsCategoryService) CreateCategory(req manageReq.MallGoodsCategoryReq) (err error) {
	if !errors.Is(
		global.GVA_DB.Where("category_level=? AND category_name=? AND is_deleted=0",
			req.CategoryLevel, req.CategoryName).First(&manage.MallGoodsCategory{}).Error,
		global.GVA_DB.Error) {
		return errors.New("该分类已存在")
	}
	rank, _ := strconv.Atoi(req.CategoryRank)
	category := manage.MallGoodsCategory{
		CategoryLevel: req.CategoryLevel,
		CategoryName:  req.CategoryName,
		CategoryRank:  rank,
		IsDeleted:     0,
		CreateTime:    common.JSONTime{Time: time.Now()},
		UpdateTime:    common.JSONTime{Time: time.Now()},
	}
	if err = utils.Verify(category, utils.GoodsCategoryVerify); err != nil {
		return errors.New(err.Error())
	}
	return global.GVA_DB.Create(&category).Error
}
