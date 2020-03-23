package repositories

import (
	"github.com/rakhmadbudiono/compare-api-framework/golang/models"
	"github.com/rakhmadbudiono/compare-api-framework/golang/utils/pg"
	 
	sq "github.com/Masterminds/squirrel"
)

type Repositories struct{}

func NewRepo() *Repositories {
	repo := &Repositories{}

	return repo
}

func (repos *Repositories) CreateBook(b *models.Book) {
	query := sq.Insert("book").
		Columns("title").
		Values(b.Title).
		RunWith(pg.Options()).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil {
		panic(err)
	}
}

func (repos *Repositories) GetBooks() []models.Book {
	sqlStatement := sq.Select("*").From("book")

	rows, err := sqlStatement.RunWith(pg.Options()).Query()

	books := []models.Book{}

	for rows.Next() {
		var b models.Book
		err = rows.Scan(&b.ID, &b.Title)
		if err != nil {
			panic(err)
		}

		books = append(books, b)
	}

	return books
}

func (repos *Repositories) GetBookByID(ID int64) *models.Book {
	query := sq.Select("*").
		From("book").
		Where(sq.Eq{"id": ID}).
		RunWith(pg.Options()).
		PlaceholderFormat(sq.Dollar)

	b := &models.Book{ID: ID}

	rows, err := query.Query()
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		rows.Scan(b.Title)
		return b
	}

	return nil
}

func (repos *Repositories) UpdateBook(b *models.Book) {
	query := sq.Update("book").
		Set("title", b.Title).
		Where(sq.Eq{"id": b.ID}).
		RunWith(pg.Options()).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil {
		panic(err)
	}
}

func (repos *Repositories) DeleteBook(ID int64) {
	query := sq.Delete("book").
		Where(sq.Eq{"id": ID}).
		RunWith(pg.Options()).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil {
		panic(err)
	}
}
