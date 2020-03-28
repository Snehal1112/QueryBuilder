package alter

import (
	"database/sql"
	"log"
)

// Table struct
type Table struct{
	name string
	db *sql.DB
}

// NewTable constructor for the Table struct.
func NewTable(name string, db *sql.DB) TableService {
	return &Table{name:name, db:db}
}

// Fields function use to manage the fields which is going to alter in table.
func (ct *Table) Fields() *Table {
	log.Println("Alter Fields called")
	return ct
}

// PrepareQuery function
func (ct *Table) prepareQuery() string {
	log.Println("Alter PrepareQuery called")
	return "ct"
}

// Execute function
func (ct *Table) Execute() (sql.Result, error) {
	log.Println("Alter execute called")
	return nil, nil
}
