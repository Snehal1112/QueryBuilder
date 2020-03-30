package datatype

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

// IsSupportLength function used to check data type support the length.
func IsSupportLength(dataType int) bool {
	for _, v := range []int{VARCHAR, BIT, CHAR, DECIMAL, INT} {
		if dataType == v {
			return true
		}
	}
	return false
}

// Get function returns the data type.
func Get(dataTypeID int) string {
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
