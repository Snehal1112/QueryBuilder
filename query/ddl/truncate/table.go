package truncate

import (
	"database/sql"
	"fmt"
)

// Table struct
type Table struct {
	name string
	db   *sql.DB
}

// NewTable constructor for the Table struct.
func TableTruncate(name string, db *sql.DB) Service {
	return &Table{name: name, db: db}
}

// PrepareQuery function
func (t *Table) prepareQuery() string {
	return fmt.Sprintf("TRUNCATE TABLE %s;", t.name)
}

// Execute function
func (t *Table) Execute() (sql.Result, error) {
	stmt, err := t.db.Prepare(t.prepareQuery())
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return stmt.Exec()
}
