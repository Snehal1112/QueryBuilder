package query

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// Database struct contains the sql DB instance.
type Database struct {
	DB *sql.DB
}

// SQLBuilder function used create the connect with database.
func SQLBuilder(driver string) *Database {
	db, err := sql.Open(driver, dataSourceName(driver))

	if err != nil {
		log.Println(err)
	}
	return &Database{db}
}

func dataSourceName(driver string) string {
	if len(driver) != 0 && driver == "sqlite3" {
		return os.Getenv("SQLITE_DB")
	}

	return fmt.Sprintf("%s:%s@/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
}

// Insert function used to perform the Insert query.
func (d *Database) Insert(table string) *InsertQuery {
	return &InsertQuery{db: d, table: table}
}

// Create function used to perform the create query.
func (d *Database) Create(table string) *CreateQuery {
	return &CreateQuery{db: d, table: table}
}

// Exec function execute the query.
func (d *Database) Exec(queryType int, query string, args ...interface{}) (sql.Result, error) {
	if queryType == DatabaseQuery {
		stmt, err := d.DB.Prepare(query)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		return stmt.Exec()
	}
	return d.DB.Exec(query, args...)
}
