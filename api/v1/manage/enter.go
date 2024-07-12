package manage

import "go-learn/service"

type ManageGroup struct {
	ManageGoodsCategoryApi
	ManageGoodsInfoApi
}

var manageGoodsCategoryService = service.ServiceGroupApp.ManageServiceGroup.ManageGoodsCategoryService
var manageGoodsInfoService = service.ServiceGroupApp.ManageServiceGroup.ManageGoodsInfoService
