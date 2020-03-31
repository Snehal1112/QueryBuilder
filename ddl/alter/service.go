package alter

import "database/sql"

type Service interface {
	prepareQuery() string
	Execute() (sql.Result, error)
}
type TableService interface {
	Rename() RenameItem
	Add() AddNewColumn
	Execute() (sql.Result, error)
}

type AddNewColumn interface {
	Column(name string, fieldType int, length interface{}, constrains []int, options ...interface{}) *AddColumn
	Service
}

type RenameItem interface {
	Column(name, newName string, fieldType int, length interface{}, constrains []int, options ...interface{}) *Rename
	Table(name, newName string) *Rename
	Service
}

// Reference https://www.mysqltutorial.org/mysql-alter-table.aspx
// Alter table
// - Rename table
// - Drop columns
// - Rename column
// - Modify columns
// - Add new columns
