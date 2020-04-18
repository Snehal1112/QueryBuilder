package ddl

import "database/sql"

type Builder struct {
	DB *sql.DB
}

func NewBuilder(db *sql.DB) Service {
	return &Builder{db}
}

func (b *Builder) Drop() DropService {
	return NewDropQuery(b)
}

func (b *Builder) Alter() AlterService {
	return NewAlterQuery(b)
}

func (b *Builder) Create() CreateService {
	return NewCreateQuery(b)
}

func (b *Builder) Truncate() TruncateService {
	return NewTruncateQuery(b)
}