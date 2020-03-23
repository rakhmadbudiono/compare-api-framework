package book

import "github.com/rakhmadbudiono/compare-api-framework/golang/util/pg"
import sq "github.com/Masterminds/squirrel"

// DAO or Data Access Object provide interface to save and fetch data into and from database
type DAO struct{}

func (dao *DAO) insert(b *Book) {
	query := sq.Insert("book").
		Columns("title").
		Values(b.Title).
		RunWith(pg.Writer()).
		PlaceholderFormat(sq.Dollar)

	result, err := query.Exec()
	if err != nil {
		panic(err)
	}

	b.ID, err = result.LastInsertId()
	if err != nil {
		panic(err)
	}
}

func (dao *DAO) getBooks() []Book {
	sqlStatement := sq.Select("*").From("book")

	rows, err := sqlStatement.RunWith(pg.Reader()).Query()

	books := []Book{}

	for rows.Next() {
		var b Book
		err = rows.Scan(&b.ID, &b.Title)
		if err != nil {
			panic(err)
		}

		books = append(books, b)
	}

	return books
}

func (dao *DAO) getBookByID(ID int64) *Book {
	query := sq.Select("title").
		From("book").
		Where(sq.Eq{"id_book": ID}).
		RunWith(pg.Reader()).
		PlaceholderFormat(sq.Dollar)

	b := &Book{ID: ID}

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

func (dao *DAO) updateBook(b *Book) {
	query := sq.Update("book").
		Set("title", b.Title).
		Where(sq.Eq{"id_book": b.ID}).
		RunWith(pg.Writer()).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil {
		panic(err)
	}
}

func (dao *DAO) deleteBook(ID int64) {
	query := sq.Delete("book").
		Where(sq.Eq{"id_book": ID}).
		RunWith(pg.Writer()).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil {
		panic(err)
	}
}
