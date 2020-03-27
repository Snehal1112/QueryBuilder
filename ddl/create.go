package ddl

import (
	"github.com/Snehal1112/QueryBuilder/ddl/create"
)

type CreateQuery struct {
	builder *Builder
}

func NewCreateQuery(b *Builder) CreateService {
	return &CreateQuery{builder: b}
}

func (c *CreateQuery) Table(name string) create.TableService {
	return create.NewTable(name, c.builder.DB)
}

func (c *CreateQuery) Database(name string) create.DatabaseService {
	return create.NewDatabase(name, c.builder.DB)
}