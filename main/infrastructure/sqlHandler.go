package infrastructure

import (
	"fmt"
	"log"
	"notes/main/domain"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	if err := godotenv.Load(); err != nil {
		log.Println("Предупреждение: Файл .env не найден, используются системные переменные окружения")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || dbName == "" {
		log.Fatal("Критическая ошибка: Переменные окружения для БД (DB_HOST, DB_USER, DB_NAME) не заданы. Проверь файл .env")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database" + err.Error())
	}

	sqlHandler := &SqlHandler{db: db}

	if err := db.AutoMigrate(&domain.Note{}); err != nil {
		log.Fatal("Failed to migrate db table" + err.Error())
	}

	return sqlHandler
}

func (handler *SqlHandler) Create(note *domain.Note) error {
	return handler.db.Create(note).Error
}

func (handler *SqlHandler) FindAll() ([]domain.Note, error) {
	var notes []domain.Note
	err := handler.db.Find(&notes).Error
	return notes, err
}

func (handler *SqlHandler) DeleteById(id string) error {
	return handler.db.Delete(&domain.Note{}, id).Error
}
