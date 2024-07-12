package service

import (
	"go-learn/service/mall"
	"go-learn/service/manage"
)

type ServiceGroup struct {
	MallServiceGroup   mall.MallServiceGroup
	ManageServiceGroup manage.ManageServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
