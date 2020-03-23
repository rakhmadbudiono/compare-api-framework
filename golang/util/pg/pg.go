package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Option can be used to configure pg connection
type Option struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

var reader *sql.DB
var writer *sql.DB

// Reader will return pg connection which has priviledges for reading data
func Reader() *sql.DB {
	return reader
}

// Writer will return pg connection which has priviledges for manipilating data
func Writer() *sql.DB {
	return writer
}

// SetupDatabase will prepare pg connection
func SetupDatabase(readerConfig Option, writerConfig Option) error {
	var err error

	reader, err = createConnection(readerConfig)
	if err != nil {
		return err
	}

	writer, err = createConnection(writerConfig)

	return err
}

func createConnection(config Option) (*sql.DB, error) {
	var db *sql.DB
	var err error

	auth := "user=" + config.User + " password=" + config.Password
	uri := " host=" + config.Host + " port=" + config.Port
	dsn :=  auth + uri + " database=" + config.Database + " sslmode=disable"

	fmt.Println(dsn)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)

	err = db.Ping()

	fmt.Println(err)

	return db, nil
}
