package initialize

import (
	"fmt"
	"time"

	"github.com/BaoTo12/go-ecommerce/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
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
	// MigrateTable()
	genTableDAO()
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

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	g.GenerateModel("go_crm_user")
	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}

// func MigrateTable() {
// 	err := global.Mdb.AutoMigrate(
// 		&po.User{},
// 		&po.Role{},
// 	)
// 	if err != nil {
// 		fmt.Println("Migration tables have errors: ", err)
// 	}
// }
