package constrain

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
func Get(constrainID int) string {
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