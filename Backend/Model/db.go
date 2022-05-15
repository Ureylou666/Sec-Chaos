package Model

// 用于配置连接数据库
import (
	"Backend/Utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	//dsn := Utils.DbUser + ":" + Utils.DbPassword + "@tcp(" + Utils.DbHost + ":" + Utils.DbPort + ")/" + Utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=" + Utils.DbHost + " user=" + Utils.DbUser + " password=" + Utils.DbPassword + " dbname=" + Utils.DbName + " port=" + Utils.DbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}

	db.AutoMigrate(&User{}, &Category{}, &Article{})
	// 连接池配置
	//sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	//sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。不要设置大于框架timeout时间
	//sqlDB.SetConnMaxLifetime(time.Second * 10)
	// 关闭数据库
	//sqlDB.Close()
}
