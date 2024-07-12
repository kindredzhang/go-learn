package manage

import (
	"github.com/gin-gonic/gin"
	"go-learn/global"
	"go-learn/model/common/response"
	manageReq "go-learn/model/manage/request"
	"go.uber.org/zap"
)

type ManageGoodsInfoApi struct {
}

func (m *ManageGoodsInfoApi) CreateGoodsInfo(c *gin.Context) {
	var mallGoodsInfo manageReq.GoodsInfoAddParam
	_ = c.ShouldBindJSON(&mallGoodsInfo)
	if err := manageGoodsInfoService.CreateMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败!"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
