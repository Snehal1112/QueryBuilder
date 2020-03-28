package drop

import (
	"database/sql"
	"fmt"
)

type Database struct {
	name string
	db *sql.DB
}

func NewDatabase(name string , db *sql.DB) DatabaseService {
	return &Database{name:name, db:db}
}

func (d *Database)prepareQuery() string {
	return fmt.Sprintf("DROP DATABASE IF EXISTS %s;", d.name)
}

func (d *Database)Execute() (sql.Result, error) {
	stmt, err := d.db.Prepare(d.prepareQuery())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec()
}