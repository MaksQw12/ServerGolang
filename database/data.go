package database

import (
	"database/sql"
	"log"
	"servergolang/models"

	"gorm.io/driver/sqlite" // GORM драйвер для SQLite
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "modernc.org/sqlite" // Импортируем драйвер SQLite
)

var DB *gorm.DB

func InitDb() {
    // Подключение через стандартный интерфейс database/sql с modernc.org/sqlite
	sqlDB, err := sql.Open("sqlite", "file:database.db?_busy_timeout=5000")
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	// Теперь используем это sqlDB для GORM
	DB, err = gorm.Open(sqlite.Dialector{
		Conn: sqlDB, // Передаем соединение к базе данных
	}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Логи для отладки
	})

	if err != nil {
		log.Fatal("Ошибка инициализации GORM с modernc.org/sqlite: ", err)
	}

	// Миграция моделей
	DB.AutoMigrate(&models.Product{}, &models.Category{})
}
