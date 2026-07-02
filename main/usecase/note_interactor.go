package usecase

import "notes/main/domain"

type NoteInteractor struct {
	NoteRepository NoteRepository
}

func (interactor *NoteInteractor) Add(note *domain.Note) error {
	return interactor.NoteRepository.Create(note)
}

func (interactor *NoteInteractor) FindAll() ([]domain.Note, error) {
	return interactor.NoteRepository.FindAll()
}

func (interactor *NoteInteractor) DeleteById(id string) error {
	return interactor.NoteRepository.DeleteById(id)
}
