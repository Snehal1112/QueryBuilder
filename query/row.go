package query

type Row []FieldService

func NewRow() *Row {
	return new(Row)
}

func (r *Row) SetField(key string, value interface{}) *Row{
	*r = append(*r, NewField(key, value))
	return r
}

func (r *Row) GetNames() (fields []string) {
	for _, field := range *r {
		fields = append(fields, field.GetKey())
	}
	return
}

func (r *Row) Transpile() (fields , placeholders []string, args []interface{}) {
	for _, field := range *r {
		fields = append(fields, field.GetKey())
		args = append(args, field.GetValue())
		placeholders = append(placeholders, "?")
	}
	return
}