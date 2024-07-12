package mall

import "go-learn/service"

type MallGroup struct {
	MallUserApi
}

var mallUserService = service.ServiceGroupApp.MallServiceGroup.MallUserService
