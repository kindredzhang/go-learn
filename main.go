package main

import (
	"go-learn/core"
	"go-learn/global"
	"go-learn/initialize"
)

func main() {

	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()

}
