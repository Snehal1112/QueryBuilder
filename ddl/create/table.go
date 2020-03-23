package create

import (
	"database/sql"
	"log"
)

type Table struct {
	name string
	db *sql.DB
}

func NewTable(name string, db *sql.DB) *Table {
	return &Table{db:db, name:name}
}

func (t *Table)Fields() *Table {
	log.Println("Fields called")
	return t
}

func (t *Table)PrepareQuery() *Table {
	log.Println("PrepareQuery called")
	return t
}

func (t *Table) Execute(query string, args ...interface{}) (sql.Result, error) {
	return t.execute(query,args)
}

func (t *Table) execute(query string, args ...interface{}) (sql.Result, error) {
	return t.db.Exec(query, args...)
}
