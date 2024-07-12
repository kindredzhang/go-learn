package mall

import (
	"github.com/gin-gonic/gin"
	v1 "go-learn/api/v1"
)

type MallUserRouter struct {
}

func (m *MallUserRouter) InitMallUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("v1")
	var mallUserApi = v1.ApiGroupApp.MallApiGroup.MallUserApi
	{
		userRouter.POST("/user/register", mallUserApi.UserRegister) //用户注册
	}
}
