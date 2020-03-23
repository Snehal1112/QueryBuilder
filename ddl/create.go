package ddl

import (
	"github.com/Snehal1112/QueryBuilder/ddl/create"
)

type CreateQuery struct {
	builder *Builder
}

func NewCreateQuery(b *Builder) *CreateQuery {
	return &CreateQuery{builder: b}
}

func (d *CreateQuery) Table(name string) *create.Table {
	return create.NewTable(name, d.builder.DB)
}

func (d *CreateQuery) Database() *create.Database {
	return create.NewDatabase()
}