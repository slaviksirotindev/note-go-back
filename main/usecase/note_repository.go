package usecase

import "notes/main/domain"

type NoteRepository interface {
	Create(note *domain.Note) error
	FindAll() ([]domain.Note, error)
	DeleteById(id string) error
}
