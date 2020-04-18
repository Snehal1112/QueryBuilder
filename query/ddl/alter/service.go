package alter

import (
	"database/sql"

	"github.com/Snehal1112/QueryBuilder/query"
)

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
	Column(name string, fieldType *query.DataType, constrains *query.Constrain) *AddColumn
	Service
}

type RenameItem interface {
	Column(name, newName string, fieldType *query.DataType, constrains *query.Constrain, options ...interface{}) *Rename
	Table(newName string) *Rename
	Service
}

// Reference https://www.mysqltutorial.org/mysql-alter-table.aspx
// Alter table
// - Rename table
// - Drop columns
// - Rename column
// - Modify columns
// - Add new columns
