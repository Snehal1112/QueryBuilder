package query

import (
	"database/sql"

	"github.com/Snehal1112/QueryBuilder/constrain"
)

type CreateDatabase struct {
	db     *Database
	name  string
}

func NewCreateDatabase(db *Database, name string) *CreateDatabase {
	return &CreateDatabase{db: db, name: name}
}

func (d *CreateDatabase) prepareQuery() string {
	return ""
}

func (d *CreateDatabase) Execute() (sql.Result, error) {
	return d.db.Exec(constrain.TableQuery, d.prepareQuery(), nil)
}