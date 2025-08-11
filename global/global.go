package global

import (
	"github.com/BaoTo12/go-ecommerce/pkg/logger"
	"github.com/BaoTo12/go-ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

/*
	Config contains
	- Redis
	- Mysql
	- ...
*/
