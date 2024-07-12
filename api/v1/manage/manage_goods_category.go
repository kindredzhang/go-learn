package manage

import (
	"github.com/gin-gonic/gin"
	"go-learn/global"
	"go-learn/model/common/response"
	manageReq "go-learn/model/manage/request"
	"go.uber.org/zap"
)

type ManageGoodsCategoryApi struct {
}

func (g *ManageGoodsCategoryApi) CreateCategory(c *gin.Context) {
	var category manageReq.MallGoodsCategoryReq
	_ = c.ShouldBindJSON(category)
	if err := manageGoodsCategoryService.CreateCategory(category); err != nil {
		global.GVA_LOG.Error("添加失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}
