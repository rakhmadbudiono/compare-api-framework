package models

type Book struct {
	ID   	int64		`json:"id_book"`
	Title 	string    	`json:"title"`
}

func NewBook(title string) *Book {
	book := &Book{
		Title: title,
	}

	return book
}

func (b *Book) GetID() int64 {
	return b.ID
}

func (b *Book) GetTitle() string {
	return b.Title
}