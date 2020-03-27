package ddl

import "database/sql"

type Builder struct {
	DB *sql.DB
}

func NewBuilder(db *sql.DB) Service {
	return &Builder{db}
}

func (b *Builder) Drop() *DropQuery {
	return NewDropQuery(b)
}

func (b *Builder) Alter() *AlterQuery {
	return NewAlterQuery()
}

func (b *Builder) Create() CreateService {
	return NewCreateQuery(b)
}