package System

import (
	"System/models"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var Connection *sql.DB

func makeConnection() *gorm.DB {
	newLogger := logger.New(

		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             500 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Warn,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,                  // Disable color
		},
	)
	newLogger = newLogger.LogMode(logger.Info)
	dsn := "root:@tcp(127.0.0.1:3306)/system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("Error connect to database", err.Error())
	}
	return db
}

func returnConnection(db *gorm.DB) *sql.DB {
	connection, err := db.DB()
	if err != nil {
		fmt.Println("Error on get database connection", err.Error())
	}
	return connection
}

func ConnectToDataBase() {
	DB = makeConnection()
	Connection = returnConnection(DB)
}

func Migrate() {
	DB.AutoMigrate(&models.Course{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.StudentCourse{})
}
