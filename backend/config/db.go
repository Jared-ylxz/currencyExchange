package config

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	db_user := AppConfig.Database.User
	db_password := AppConfig.Database.Password
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/currency_exchange?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(AppConfig.Database.ConnMaxLifetime) * time.Second) // 10 seconds

	if err := db.AutoMigrate(&models.User{}, &models.ExchangeRate{}, &models.Article{}); err != nil {
		panic("failed to migrate database")
	}

	global.Db = db

}
