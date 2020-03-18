package query

import (
	"database/sql"
	"fmt"
)

type CreateDatabase struct {
	db     *Database
	name  string
}

func NewCreateDatabase(db *Database, name string) *CreateDatabase {
	return &CreateDatabase{db: db, name: name}
}

func (d *CreateDatabase) prepareQuery() string {
	// TODO: validate the database name.
	return fmt.Sprintf("CREATE DATABASE %s", d.name)
}

func (d *CreateDatabase) Execute() (sql.Result, error) {
	return d.db.Exec(DatabaseQuery, d.prepareQuery(), nil)
}