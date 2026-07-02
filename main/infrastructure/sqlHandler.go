package infrastructure

import (
	"notes/main/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	dsn := "host=localhost user=postgres password=20021216bur dbname=note port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.Note{})
	return &SqlHandler{db: db}
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
