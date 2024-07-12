package global

import (
	"github.com/spf13/viper"
	"go-learn/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	//"go-learn/config"
)

var (
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
)
