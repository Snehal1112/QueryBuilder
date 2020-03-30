package create

import (
	"database/sql"
	"fmt"
)

type Database struct {
	db   *sql.DB
	name string
}

func NewDatabase(name string, db *sql.DB) DatabaseService {
	return &Database{name: name, db: db}
}

func (d *Database) prepareQuery() string {
	return fmt.Sprintf("CREATE DATABASE IF NOTE EXISTS %s;", d.name)
}

func (d *Database) Execute() (sql.Result, error) {
	// TODO: Code duplication with Table.Execute function
	//  move this code to create.go which is common struct for the
	//  create table as well as database.
	stmt, err := d.db.Prepare(d.prepareQuery())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec()
}
