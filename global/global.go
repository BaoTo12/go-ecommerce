package global

import (
	"github.com/BaoTo12/go-ecommerce/pkg/logger"
	"github.com/BaoTo12/go-ecommerce/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)

/*
	Config contains
	- Redis
	- Mysql
	- ...
*/
