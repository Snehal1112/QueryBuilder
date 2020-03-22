package ddl

import "github.com/Snehal1112/QueryBuilder/ddl/create"

type CreateQuery struct {}

func NewCreateQuery() *CreateQuery {
	return &CreateQuery{}
}

func (D *CreateQuery) Table() *create.Table {
	return create.NewTable()
}

func (D *CreateQuery) Database() *create.Database {
	return create.NewDatabase()
}