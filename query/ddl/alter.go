package ddl

import "github.com/Snehal1112/QueryBuilder/query/ddl/alter"

type AlterQuery struct {
	builder *Builder
}

func NewAlterQuery(builder *Builder) AlterService {
	return &AlterQuery{builder: builder}
}

func (a *AlterQuery) Table(name string) alter.TableService {
	return alter.NewTable(name, a.builder.DB)
}
