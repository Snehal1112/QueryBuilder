package query

import (
	"database/sql"
	"fmt"
	"strings"
)

type DropTableColumns struct {
	db   *Database
	name string
	columns []interface{}
}

func NewDropTableColumns(db *Database, name string) *DropTableColumns {
	return &DropTableColumns{db: db, name: name}
}

func (d *DropTableColumns) Field(column string) *DropTableColumns {
	d.columns = append(d.columns, column)
	return d
}

func (d *DropTableColumns) Fields(column []interface{}) *DropTableColumns {
	d.columns = append(d.columns, column...)
	return d
}

func (d *DropTableColumns) Execute() (sql.Result, error) {
	return d.db.Exec(DatabaseQuery, d.prepareQuery(), nil)
}

func (d *DropTableColumns) prepareQuery() string {
	var dropColumns []string
	for _, column := range d.columns {
		dropColumns = append(dropColumns, fmt.Sprintf("DROP COLUMN %v", column))
	}
	return fmt.Sprintf("ALTER TABLE %s %s;", d.name, strings.Join(dropColumns, ", "))
}


