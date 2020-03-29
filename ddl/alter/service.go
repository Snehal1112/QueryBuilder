package alter

import "database/sql"

type Service interface {
	prepareQuery() string
	Execute() (sql.Result, error)
}
type TableService interface {
	Rename(newName string) *Table
	Add() AddNewColumn
	Execute() (sql.Result, error)
}

type AddNewColumn interface {
	Column(name string, fieldType int, length interface{}, constrains []int, options ...interface{}) *AddColumn
	Service
}

// Reference https://www.mysqltutorial.org/mysql-alter-table.aspx
// Alter table
// - Rename table
// - Drop columns
// - Rename column
// - Modify columns
// - Add new columns