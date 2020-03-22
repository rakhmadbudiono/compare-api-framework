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

var option *sql.DB

func Options() *sql.DB {
	return option
}

// SetupDatabase will prepare pg connection
func SetupDatabase(optionConfig Option) error {
	var err error

	option, err = createConnection(optionConfig)
	if err != nil {
		return err
	}

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
