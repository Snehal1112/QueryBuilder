package query

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(driver string, url string) *Database {
	db, err := sql.Open(driver, url)
	if err != nil {
		logrus.Println(err)
	}

	return &Database{db}
}

func (d *Database)Insert(table string) *InsertQuery {
	return &InsertQuery{db: d, table: table}
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error){
	return d.DB.Exec(query, args...)
}