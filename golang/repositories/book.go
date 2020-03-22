package repositories

import (
	"github.com/rakhmadbudiono/compare-api-framework/golang/models"
	"github.com/rakhmadbudiono/compare-api-framework/golang/utils/pg"
	 
	sq "github.com/Masterminds/squirrel"
)

type Repositories struct{}

func (repos *Repositories) createBook(b *Book) {
	query := sq.Insert("book").
		Columns("title").
		Values(b.Title).
		RunWith(pg.Options()).
		PlaceholderFormat(sq.Dollar)

	result, err := query.Exec()
	if err != nil {
		panic(err)
	}
}

func (repos *Repositories) getBooks() []Book {
	sqlStatement := sq.Select("*").From("book")

	rows, err := sqlStatement.RunWith(pg.Options()).Query()

	books := []Book{}

	for rows.Next() {
		var b Book
		err = rows.Scan(&b.ID, &b.Name)
		if err != nil {
			panic(err)
		}

		books = append(books, b)
	}

	return books
}

func (repos *Repositories) getBookByID(ID int64) *Book {
	query := sq.Select("*").
		From("book").
		Where(sq.Eq{"id": ID}).
		RunWith(pg.Options()).
		PlaceholderFormat(sq.Dollar)

	b := &Book{ID: ID}

	rows, err := query.Query()
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		rows.Scan(b.Name)
		return b
	}

	return nil
}

func (repos *Repositories) updateBook(b *Book) {
	query := sq.Update("book").
		Set("name", b.Name).
		Where(sq.Eq{"id": b.ID}).
		RunWith(pg.Options()).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil {
		panic(err)
	}
}

func (repos *Repositories) deleteBook(ID int64) {
	query := sq.Delete("book").
		Where(sq.Eq{"id": ID}).
		RunWith(pg.Options()).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil {
		panic(err)
	}
}
