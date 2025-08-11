package initialize

import (
	"fmt"
	"time"

	"github.com/BaoTo12/go-ecommerce/global"
	"github.com/BaoTo12/go-ecommerce/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkError(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	mSetting := global.Config.MYSQL

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, mSetting.Username, mSetting.Password, mSetting.Host, mSetting.Port, mSetting.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	// ! Check error
	checkError(err, "Mysql Initialization Failed ~~!")
	global.Logger.Info("Initializing Mysql Successfully")
	global.Mdb = db

	// TODO: Set pool
	SetPool()

	// TODO: Migration
	MigrateTable()
}
func SetPool() {
	mSetting := global.Config.MYSQL

	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error %s \n", err)
	}
	sqlDb.SetMaxIdleConns(mSetting.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(time.Duration(mSetting.ConnMaxLifetime))
	sqlDb.SetMaxOpenConns(mSetting.MaxOpenConns)
}
func MigrateTable() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("Migration tables have errors: ", err)
	}
}
