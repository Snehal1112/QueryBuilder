package ddl

type Builder struct {}

func NewBuilder() Service {
	return &Builder{}
}

func (b Builder) Drop() *DropQuery {
	return NewDropQuery()
}

func (b Builder) Alter() *AlterQuery {
	return NewAlterQuery()
}

func (b Builder) Create() *CreateQuery {
	return NewCreateQuery()
}