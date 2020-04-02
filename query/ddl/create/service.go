package create

import (
	"database/sql"

	"github.com/Snehal1112/QueryBuilder/query"
)

type Service interface {
	prepareQuery() string
	Execute() (sql.Result, error)
}
type TableService interface {
	Field(name string, dataType *query.DataType, constrain *query.Constrain) *Table
	SetPrimaryKey(fields []string) *Table
	NewForeignKeyConstrain(constrain, foreignKey, fkTable string) *Table
	SetForeignKey(onUpdate, onDelete interface{}) *Table
	Service
}

type DatabaseService interface {
	Service
}
