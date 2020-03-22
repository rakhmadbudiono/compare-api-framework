package models

type Book struct {
	ID   	int			`json:"id_book"`
	Title 	string    	`json:"title"`
}

func NewBook(title string) *Book {
	book := &Book{
		Title: title,
	}

	return book
}

func (m *Book) GetID() int {
	return m.ID
}

func (m *Book) GetTitle() string {
	return m.Title
}