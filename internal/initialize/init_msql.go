package initialize

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql

	// Kiểm tra và sử dụng biến môi trường nếu có
	host := getEnvOrDefault("MYSQL_HOST", m.Host)
	port := getEnvOrDefaultInt("MYSQL_PORT", m.Port)
	username := getEnvOrDefault("MYSQL_USER", m.Username)
	password := getEnvOrDefault("MYSQL_PASSWORD", m.Password)
	dbname := getEnvOrDefault("MYSQL_DB", m.Dbname)

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdb = db

	// set Pool
	SetPool()
	// genTableDAO()
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

// auto create table in mysql
// refer https://gorm.io/docs/migration.html#AutoMigrate
func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&model.Customerv1{},
		&model.Customer{},
	)

	if err != nil {
		global.Logger.Error("mysql auto migrate error", zap.Error(err))
	} else {
		global.Logger.Info("mysql auto migrate success")
	}
}

// generate table in mysql to go struct
func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(global.Mdb)
	g.GenerateModel("test_gen_mysql_to_go")

	g.Execute()
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}
