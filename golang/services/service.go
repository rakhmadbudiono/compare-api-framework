package services

import (
	"github.com/rakhmadbudiono/compare-api-framework/golang/repositories"
)

type BookService struct {
	repo *Repositories
}

func (s *BookService) CreateBook(p *Book) int64 {
	s.repo.createBook(p)
	return p.ID
}

func (s *BookService) GetBooks(q *Query) []Book {
	return s.repo.getBooks(q)
}

func (s *BookService) GetBookByID(ID int64) *Book {
	return s.repo.getBookByID(ID)
}

func (s *BookService) UpdateBook(p *Book) {
	s.repo.updateBook(p)
}

func (s *BookService) DeleteBook(ID int64) {
	s.repo.deleteBook(ID)
}