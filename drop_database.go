package query

import (
	"database/sql"
	"fmt"
)

type DropDatabase struct {
	db   *Database
	name string
}

func NewDropDatabase(db *Database, name string) *DropDatabase {
	return &DropDatabase{db: db, name: name}
}

func (d *DropDatabase) Execute() (sql.Result, error) {
	return d.db.Exec(DatabaseQuery, d.prepareQuery(), nil)
}

func (d *DropDatabase) prepareQuery() string {
	return fmt.Sprintf("DROP DATABASE IF EXISTS %s", d.name)
}


