package ddl

import "github.com/Snehal1112/QueryBuilder/ddl/drop"

type DropQuery struct {}

func NewDropQuery() *DropQuery {
	return &DropQuery{}
}

func (D *DropQuery) CreateTable() *drop.Table {
	return drop.NewTable()
}

func (D *DropQuery) CreateDatabase() *drop.Database {
	return drop.NewDatabase()
}