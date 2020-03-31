package dml

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (D Builder) Update() {
	panic("implement me")
}

func (D Builder) Delete() {
	panic("implement me")
}

func (D Builder) Insert() {
	panic("implement me")
}
