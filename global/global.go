package global

import (
	"github.com/spf13/viper"
	"go-learn/config"
	//"go-learn/config"
)

var (
	//GVA_DB     *gorm.DB
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Server
)
