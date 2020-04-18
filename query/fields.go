package query

type Field struct {
	Key   string
	Value interface{}
}

func NewField(key string, value interface{}) FieldService {
	return &Field{Key: key, Value: value}
}

func (f *Field) GetKey() string {
	return f.Key
}

func (f *Field) GetValue() interface{} {
	return f.Value
}