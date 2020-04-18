package query

type RowsService interface {
	SetRow(row *Row) *Rows
	GetValues() (args []interface{})
}

type RowService interface {
	SetField(key string, value interface{}) *Row
	GetNames() (fields []string)
	Transpile() (fields , placeholders []string, args []interface{})
}

type FieldService interface {
	GetKey() string
	GetValue() interface{}
}