package dml

import "github.com/Snehal1112/QueryBuilder/query/dml/insert"

type InsertQuery struct {
	builder *Builder
}

func NewInsertQuery(builder *Builder) InsertService {
	return &InsertQuery{builder: builder}
}

func (i *InsertQuery) Into(name string) insert.Service {
	return insert.NewQueryInsert(name, i.builder.DB)
}