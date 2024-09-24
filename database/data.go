package database

import (
	"database/sql"
	"log"
	"servergolang/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func InitDb() {
   
	sqlDB, err := sql.Open("sqlite", "file:database.db?_busy_timeout=5000")
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	
	DB, err = gorm.Open(sqlite.Dialector{
		Conn: sqlDB, 
	}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Ошибка инициализации GORM с modernc.org/sqlite: ", err)
	}


	DB.AutoMigrate(&models.Product{}, &models.Category{})
}
