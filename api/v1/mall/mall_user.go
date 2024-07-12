package mall

import (
	"github.com/gin-gonic/gin"
	"go-learn/global"
	"go-learn/model/common/response"
	mallReq "go-learn/model/mall/request"
	"go-learn/utils"
	"go.uber.org/zap"
)

type MallUserApi struct {
}

func (m *MallUserApi) UserRegister(c *gin.Context) {
	var req mallReq.RegisterUserParam
	_ = c.ShouldBindJSON(req)
	if err := utils.Verify(req, utils.MallUserRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err := mallUserService.RegisterUser(req); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
	}
	response.OkWithMessage("注册成功", c)
}
