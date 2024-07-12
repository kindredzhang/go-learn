package main

import (
	"go-learn/core"
	"go-learn/global"
)

func main() {

	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()

}
