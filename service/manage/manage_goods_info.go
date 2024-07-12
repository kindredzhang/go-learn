package manage

import (
	"errors"
	"go-learn/global"
	"go-learn/model/common"
	"go-learn/model/common/enum"
	"go-learn/model/manage"
	manageReq "go-learn/model/manage/request"
	"go-learn/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type ManageGoodsInfoService struct {
}

func (m *ManageGoodsInfoService) CreateMallGoodsInfo(req manageReq.GoodsInfoAddParam) (err error) {
	var goodsCategory manage.MallGoodsCategory
	err = global.GVA_DB.Where("category_id=?  AND is_deleted=0", req.GoodsCategoryId).First(&goodsCategory).Error
	if goodsCategory.CategoryLevel != enum.LevelThree.Code() {
		return errors.New("分类数据异常")
	}
	if !errors.Is(global.GVA_DB.Where("goods_name=? AND goods_category_id=?", req.GoodsName, req.GoodsCategoryId).First(&manage.MallGoodsInfo{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在相同的商品信息")
	}
	originalPrice, _ := strconv.Atoi(req.OriginalPrice)
	sellingPrice, _ := strconv.Atoi(req.SellingPrice)
	stockNum, _ := strconv.Atoi(req.StockNum)
	goodsSellStatus, _ := strconv.Atoi(req.GoodsSellStatus)
	goodsInfo := manage.MallGoodsInfo{
		GoodsName:          req.GoodsName,
		GoodsIntro:         req.GoodsIntro,
		GoodsCategoryId:    req.GoodsCategoryId,
		GoodsCoverImg:      req.GoodsCoverImg,
		GoodsDetailContent: req.GoodsDetailContent,
		OriginalPrice:      originalPrice,
		SellingPrice:       sellingPrice,
		StockNum:           stockNum,
		Tag:                req.Tag,
		GoodsSellStatus:    goodsSellStatus,
		CreateTime:         common.JSONTime{Time: time.Now()},
		UpdateTime:         common.JSONTime{Time: time.Now()},
	}
	if err = utils.Verify(goodsInfo, utils.GoodsAddParamVerify); err != nil {
		return errors.New(err.Error())
	}
	err = global.GVA_DB.Create(&goodsInfo).Error
	return err
}
