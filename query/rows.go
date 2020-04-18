package query

type Rows []Row

func NewRows() *Rows {
	return new(Rows)
}

func (rs *Rows)SetRow(row *Row) *Rows {
	*rs = append(*rs, *row)
	return rs
}

func (rs *Rows) GetValues() (args []interface{}) {
	for _, v := range *rs {
		for _, field := range v {
			args = append(args, field.GetValue())
		}
	}
	return
}

