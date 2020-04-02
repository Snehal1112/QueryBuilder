package create

import "database/sql"

type Service interface {
	prepareQuery() string
	Execute() (sql.Result, error)
}
type TableService interface {
	Field(name string, fieldType int, length interface{}, constrains []int) *Table
	SetPrimaryKey(fields []string) *Table
	NewForeignKeyConstrain(constrain, foreignKey, fkTable string) *Table
	SetForeignKey(onUpdate, onDelete interface{}) *Table
	Service
}

type DatabaseService interface {
	Service
}
