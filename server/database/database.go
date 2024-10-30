package database

import (
	"log"
	"os"

	"github.com/njmhywn-dev/go-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	user := os.Getenv("db_user")
	name := os.Getenv("db_name")

	dsn := user + ":@tcp(127.0.0.1:3307)/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connection failed.")
	}

	log.Println("Connection successful.")

	db.AutoMigrate(new(model.Blog))
	db.AutoMigrate(new(model.User))

	DBConn = db
}
