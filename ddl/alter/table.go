package alter

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

const (
	RenameTable = iota + 1
	DropColumns
	RenameColumns
	ModifyColumns
	AddColumns
)

// Table struct
type Table struct {
	name string
	db   *sql.DB

	queryType int
	addCol    AddNewColumn
}

// NewTable constructor for the Table struct.
func NewTable(name string, db *sql.DB) TableService {
	return &Table{name: name, db: db}
}

func (t *Table) Add() AddNewColumn {
	if t.queryType != 0 {
		logrus.WithFields(logrus.Fields{
			"expectedQuery": AddColumns,
			"currentQuery":  t.queryType,
		}).Error("Already different query type is set")
	}
	t.queryType = AddColumns
	t.addCol = NewAddColumn(t)
	return t.addCol
}

func (t *Table) Rename(newName string) *Table {
	panic("implement me")
}

// PrepareQuery function
func (t *Table) prepareQuery() string {
	switch t.queryType {
	case DropColumns:
		return ""
	default:
		return t.addCol.prepareQuery()
	}
}

// Execute function
func (t *Table) Execute() (sql.Result, error) {
	t.queryType = 0
	stmt, err := t.db.Prepare(t.prepareQuery())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec()
}
