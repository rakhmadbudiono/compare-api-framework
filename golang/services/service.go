package services

import (
	"github.com/rakhmadbudiono/compare-api-framework/golang/models"
	"github.com/rakhmadbudiono/compare-api-framework/golang/repositories"
)

type BookService struct {
	repo *repositories.Repositories
}

func NewBookService() *BookService {
	service := &BookService{ repo: repositories.NewRepo() }

	return service
}

func (s *BookService) CreateBook(p *models.Book) int64 {
	s.repo.CreateBook(p)
	return p.ID
}

func (s *BookService) GetBooks() []models.Book {
	return s.repo.GetBooks()
}

func (s *BookService) GetBookByID(ID int64) *models.Book {
	return s.repo.GetBookByID(ID)
}

func (s *BookService) UpdateBook(p *models.Book) {
	s.repo.UpdateBook(p)
}

func (s *BookService) DeleteBook(ID int64) {
	s.repo.DeleteBook(ID)
}