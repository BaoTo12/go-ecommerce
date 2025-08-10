package initialize

import (
	"github.com/BaoTo12/go-ecommerce/global"
	"github.com/BaoTo12/go-ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(&global.Config.LOGGER)
}
