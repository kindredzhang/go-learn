package v1

import (
	"go-learn/api/v1/mall"
	"go-learn/api/v1/manage"
)

type ApiGroup struct {
	MallApiGroup   mall.MallGroup
	ManageApiGroup manage.ManageGroup
}

var ApiGroupApp = new(ApiGroup)
