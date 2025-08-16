package initialize

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"time"

	"github.com/BaoTo12/go-ecommerce/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func openGORMWithRetry(connectString string, timeout time.Duration, initialDelay time.Duration) error {
	deadline := time.Now().Add(timeout)
	attempt := 0
	for {
		attempt++
		// Try to open GORM
		db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
		if err == nil {
			// get underlying *sql.DB and ping to ensure accept connections
			sqlDB, err2 := db.DB()
			if err2 == nil {
				// TODO: Set pool
				mysqlSetting := global.Config.MYSQL
				sqlDB.SetMaxIdleConns(mysqlSetting.MaxIdleConns)
				sqlDB.SetConnMaxLifetime(time.Duration(mysqlSetting.ConnMaxLifetime))
				sqlDB.SetMaxOpenConns(mysqlSetting.MaxOpenConns)

				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				errPing := sqlDB.PingContext(ctx)
				cancel()
				if errPing == nil {
					// success
					global.Mdb = db
					log.Printf("connected to mysql (attempt %d)", attempt)
					return nil
				}
				_ = sqlDB.Close()
				err = fmt.Errorf("ping failed after open: %w", errPing)
			}
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("could not connect to mysql after %s: last error: %w", timeout, err)
		}

		// exponential backoff with jitter
		// sleep = min( cap, initialDelay * 2^(attempt-1) ) +- jitter
		capDelay := 5 * time.Second
		backoff := float64(initialDelay) * math.Pow(2, float64(attempt-1))
		if backoff > float64(capDelay) {
			backoff = float64(capDelay)
		}
		// jitter: +/- 30%
		jitter := (rand.Float64()*0.6 - 0.3) * backoff
		sleep := time.Duration(backoff + jitter)

		log.Printf("mysql not ready (attempt %d): %v. retrying in %v", attempt, err, sleep)
		time.Sleep(sleep)
	}
}
func InitMysql() {
	mSetting := global.Config.MYSQL

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, mSetting.Username, mSetting.Password, mSetting.Host, mSetting.Port, mSetting.Dbname)
	// try up to 60s; initial delay 500ms
	err := openGORMWithRetry(s, 60*time.Second, 500*time.Millisecond)
	// ! Check error
	if err != nil {
		global.Logger.Error("Mysql Initialization Failed ~~!", zap.Error(err))
	}

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
