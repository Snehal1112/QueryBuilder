package ddl

import "github.com/Snehal1112/QueryBuilder/ddl/drop"

type DropQuery struct {
	builder *Builder
}

func NewDropQuery(b *Builder) *DropQuery {
	return &DropQuery{builder: b}
}

func (D *DropQuery) Table() *drop.Table {
	return drop.NewTable()
}

func (D *DropQuery) Database() *drop.Database {
	return drop.NewDatabase()
}