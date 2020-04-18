package ddl

import (
	"github.com/Snehal1112/QueryBuilder/query/ddl/truncate"
)

type TruncateQuery struct {
	builder *Builder
}

func NewTruncateQuery(b *Builder) TruncateService {
	return &TruncateQuery{builder: b}
}

func (t TruncateQuery) Table(name string) truncate.Service {
	return truncate.TableTruncate(name, t.builder.DB)
}
