package query

import (
	"strings"

	"github.com/spf13/cast"
)

// Constance define the constrains.
const (
	NOTNULL = iota + 100
	UNIQUE
	PK
	FK
	CHECK
	AI
)

type Constrain struct {
	Cons []int
}

func NewConstrain(cons interface{}) *Constrain {
	return &Constrain{Cons: cast.ToIntSlice(cons)}
}

func (c *Constrain) AsString() string {
	var constrain []string
	for _, con := range c.Cons {
		constrain = append(constrain, GetConstrain(con))
	}
	return strings.Join(constrain, " ")
}

// Get helper function used to get the constance value.
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

// Constants used to mange the referential integrity between the child and parent tables
// by using the ON DELETE and ON UPDATE clauses
const (
	SetNull = iota + 2000
	Cascade
	Restrict
	NoAction
	SetDefault
)

func GetReferenceOpt(referenceOpt int) string {
	switch referenceOpt {
	case SetNull:
		return "SET NULL"
	case Cascade:
		return "CASCADE"
	case Restrict:
		return "RESTRICT"
	case NoAction:
		return "NO ACTION"
	default:
		// TODO: SET DEFAULT is not supported right now
		return "SET NULL"
	}
}
