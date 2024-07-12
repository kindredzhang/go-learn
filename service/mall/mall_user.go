package mall

import (
	"errors"
	"go-learn/global"
	"go-learn/model/common"
	"go-learn/model/mall"
	mallReq "go-learn/model/mall/request"
	"go-learn/utils"
	"gorm.io/gorm"
	"time"
)

type MallUserService struct {
}

func (s MallUserService) RegisterUser(req mallReq.RegisterUserParam) (err error) {
	if !errors.Is(global.GVA_DB.Where("login_name", req.LoginName).First(&mall.MallUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已存在")
	}

	return global.GVA_DB.Create(&mall.MallUser{
		LoginName:     req.LoginName,
		PasswordMd5:   utils.MD5V([]byte(req.Password)),
		IntroduceSign: "这个人很懒，什么都没留下",
		CreateTime:    common.JSONTime{Time: time.Now()},
	}).Error
}
