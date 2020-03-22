package ddl

import "github.com/Snehal1112/QueryBuilder/ddl/alter"

type AlterQuery struct{}

func NewAlterQuery() *AlterQuery {
	return &AlterQuery{}
}

func (D *AlterQuery) Table() *alter.Table {
	return alter.NewTable()
}
