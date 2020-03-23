package book

type Service struct {
	dao DAO
}

func (s *Service) CreateBook(b *Book) int64 {
	s.dao.insert(b)
	return b.ID
}

func (s *Service) GetBooks() []Book {
	return s.dao.getBooks()
}

func (s *Service) GetBookByID(ID int64) *Book {
	return s.dao.getBookByID(ID)
}

func (s *Service) UpdateBook(b *Book) {
	s.dao.updateBook(b)
}

func (s *Service) DeleteBook(ID int64) {
	s.dao.deleteBook(ID)
}
