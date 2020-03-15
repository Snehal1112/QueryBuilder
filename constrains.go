package query

// Constance define the constrains.
const (
	NOTNULL = iota + 100
	UNIQUE
	PK
	FK
	CHECK
	AI
)

// GetConstrain helper function used to get the constance value.
func GetConstrain(constrainID int) string {
	switch constrainID {
	case NOTNULL:
		return "NOT NULL"
	case UNIQUE:
		return "UNIQUE"
	case PK:
		return "PRIMARY KEY"
	case FK:
		return "FOREIGN KEY"
	case CHECK:
		return "CHECK"
	case AI:
		return "AUTO_INCREMENT"
	}
	return "NULL"
}

// Constance define the type of the query either it can be
// database/table level query.
const (
	DatabaseQuery = iota + 0
	TableQuery
)

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

// GetDataType function returns the data type.
func GetDataType(dataTypeID int) string {
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

// Constants used to mange the referential integrity between the child and parent tables
// by using the ON DELETE and ON UPDATE clauses
const (
	SETNULL = iota + 2000
	CASCADE
	RESTRICT
	NOACTION
	SETDEFAULT
)

func GetReferenceOpt(referenceOpt int) string {
	switch referenceOpt {
	case SETNULL:
		return "SET NULL"
	case CASCADE:
		return "CASCADE"
	case RESTRICT:
		return "RESTRICT"
	case NOACTION:
		return "NO ACTION"
	case SETDEFAULT:
		// TODO: SET DEFAULT is not supported right now
		return "SET NULL"
	}
	return "SET NULL"
}