package manage

import (
	"github.com/gin-gonic/gin"
	v1 "go-learn/api/v1"
)

type ManageGoodsInfoRouter struct {
}

func (m *ManageGoodsInfoRouter) InitManageGoodsInfoRouter(Router *gin.RouterGroup) {
	mallGoodsInfoRouter := Router.Group("v1")
	var mallGoodsInfoApi = v1.ApiGroupApp.ManageApiGroup.ManageGoodsInfoApi
	{
		mallGoodsInfoRouter.POST("goods", mallGoodsInfoApi.CreateGoodsInfo)
	}
}
