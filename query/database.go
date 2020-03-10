package query

import (
	"database/sql"

	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(driver string) *Database {
	db, err := sql.Open(driver, dataSourceName(driver))

	if err != nil {
		log.Println(err)
	}
	return &Database{db}
}

func dataSourceName(driver string) string {
	if len(driver) != 0  && driver == "sqlite3" {
		return os.Getenv("SQLITE_DB")
	}

	return  fmt.Sprintf("%s:%s@/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
}

func (d *Database)Insert(table string) *InsertQuery {
	return &InsertQuery{db: d, table: table}
}

func (d *Database)Create(table string) *CreateQuery {
	return &CreateQuery{db: d, table:table}
}

func (d *Database) Exec(queryType int, query string, args ...interface{}) (sql.Result, error){

	if queryType == DatabaseQuery {
		stmt, err := d.DB.Prepare(query)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		return stmt.Exec()
	} else {
		return d.DB.Exec(query, args...)
	}
}