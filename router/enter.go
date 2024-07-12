package router

import (
	"go-learn/router/mall"
	"go-learn/router/manage"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	Mall   mall.MallRouterGroup
}

var RouterGroupApp = new(RouterGroup)
