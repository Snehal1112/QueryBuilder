package dml

import "database/sql"

type Builder struct{
	DB *sql.DB
}

func NewBuilder(db *sql.DB) Service {
	return &Builder{db}
}

func (D *Builder) Update() {
	panic("implement me")
}

func (D *Builder) Delete() {
	panic("implement me")
}

func (D *Builder) Insert() InsertService {
	return NewInsertQuery(D)
}
