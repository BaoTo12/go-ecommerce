package initialize

import (
	"fmt"

	"github.com/BaoTo12/go-ecommerce/global"
	"gorm.io/gorm"
)

func InitMysql() {
	mSetting := global.Config.MYSQL

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	string := fmt.Sprintf(dsn, mSetting.Username, mSetting.Password, mSetting.Host, mSetting.Port, mSetting.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
