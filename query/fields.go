package query

type Row map[string]interface{}

func NewRow() *Row {
	fields := make(Row)
	return &fields
}

func (f *Row) GetNames() (fields []string) {
	for k, _ := range *f {
		fields = append(fields, k)
	}
	return
}

func (f *Row) Transpile() (fields , placeholders []string, args []interface{}) {
	for k, v := range *f {
		fields = append(fields, k)
		args = append(args, v)
		placeholders = append(placeholders, "?")
	}
	return
}

func (f *Row) Placeholders() (placeholders []string) {
	for i:=0; i < len(*f); i++ {
		placeholders = append(placeholders, "?")
	}

	return
}

func (f *Row) SetField(key string, value interface{}) *Row {
	(*f)[key] = value
	return f
}

type Rows []*Row

func NewRows() *Rows {
	return &Rows{}
}

func (r *Rows) SetField(row *Row) Rows{
	*r = append(*r, row)
	return *r
}

func (r *Rows) GetValues() (args []interface{}) {
	for _, v := range *r {
		for _, i := range *v {
			args = append(args, i)
		}
	}
	return
}

