package query

import (
	"database/sql"
	"fmt"
	"strings"
)

type DropTable struct {
	db   *Database
	name []string
}

func NewDropTable(db *Database, name []string) *DropTable {
	return &DropTable{db: db, name: name}
}

func (d *DropTable) Execute() (sql.Result, error) {
	return d.db.Exec(DatabaseQuery, d.prepareQuery(), nil)
}

func (d *DropTable) prepareQuery() string {
	return fmt.Sprintf("DROP TABLE %s", strings.Join(d.name, ", "))
}


