package query

import "github.com/spf13/cast"

const (
	VARCHAR = iota + 1000
	TIMESTAMP
	BIT
	BOOLEAN
	CHAR
	DATE
	DATETIME
	DECIMAL
	ENUM
	INT
	JSON
	TIME
)

type DataType struct {
	dataType int
	length interface{}
}

func NewDataType(dataType int, length interface{}) *DataType {
	return &DataType{dataType: dataType,length:length}
}

func (d *DataType) AsString() string {
	fieldDataType := d.Get(d.dataType)
	if d.isSupportLength(d.dataType) {
		fieldDataType += "(" + cast.ToString(d.length) + ")"
	}
	return fieldDataType
}
// isSupportLength function used to check data type support the length.
func (d *DataType)isSupportLength(dataType int) bool {
	for _, v := range []int{VARCHAR, BIT, CHAR, DECIMAL, INT} {
		if dataType == v {
			return true
		}
	}
	return false
}

// Get function returns the data type.
func (d *DataType) Get(dataTypeID int) string {
	switch dataTypeID {
	case VARCHAR:
		return "VARCHAR"
	case TIMESTAMP:
		return "TIMESTAMP"
	case BIT:
		return "BIT"
	case BOOLEAN:
		return "BOOLEAN"
	case CHAR:
		return "CHAR"
	case DATE:
		return "DATE"
	case DATETIME:
		return "DATETIME"
	case DECIMAL:
		return "DECIMAL"
	case ENUM:
		return "ENUM"
	case INT:
		return "INT"
	case JSON:
		return "JSON"
	case TIME:
		return "TIME"
	}
	return "VARCHAR"
}
