package ddl

import "github.com/Snehal1112/QueryBuilder/query/ddl/drop"

type DropQuery struct {
	builder *Builder
}

func NewDropQuery(b *Builder) DropService {
	return &DropQuery{builder: b}
}

func (d *DropQuery) Table(name []string) drop.TableService {
	return drop.NewTable(name, d.builder.DB)
}

func (d *DropQuery) Database(name string) drop.DatabaseService {
	return drop.NewDatabase(name, d.builder.DB)
}
