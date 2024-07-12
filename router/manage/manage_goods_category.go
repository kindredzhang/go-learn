package manage

import (
	"github.com/gin-gonic/gin"
	v1 "go-learn/api/v1"
)

type ManageGoodsCategoryRouter struct {
}

func (r *ManageGoodsCategoryRouter) InitManageGoodsCategoryRouter(Router *gin.RouterGroup) {
	goodsCategoryRouter := Router.Group("v1")

	var goodsCategoryApi = v1.ApiGroupApp.ManageApiGroup.ManageGoodsCategoryApi
	{
		goodsCategoryRouter.POST("categories", goodsCategoryApi.CreateCategory)
	}
}
