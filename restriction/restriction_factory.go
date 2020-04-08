package restriction

import (
	"strings"
	"sync"
)

// Operators
const (

	EQ = iota + 1
	NE
	LE
	GE
	LT
	GT
	AND
	OR
	BETWEEN
	LINK
	IN
	ISNULL
)

type Operator int
func (o Operator)ToString() string {
	switch o {
	case EQ:
		return "="
	case NE:
		return "!="
	case LE:
		return "<="
	case GE:
		return ">="
	case LT:
		return "<"
	case GT:
		return ">"
	case OR:
		return "OR"
	case AND:
		return "AND"
	default:
		return ""
	}
}

type ConditionStack struct {
	toString func()string
	operator Operator
}

type Restriction struct {
	sync.Mutex
	stack []*ConditionStack
	restrictionToString string
}

func (r *Restriction) Transpile() string {
	var condition []string
	for _, f := range r.stack {
		condition = append(condition, f.toString(), f.operator.ToString())
	}
	return strings.Join(condition, " ")
}

func (r *Restriction) AddCondition(condition func() string, operator Operator) {
	r.Lock()
	defer r.Unlock()
	r.stack = append(r.stack, &ConditionStack{toString:condition, operator:operator})
}

func (r *Restriction) Factory(restrictionType int) ConditionService {
	switch restrictionType {
	case AND:
		return NewAnd()
	case OR:
		return NewOr()
	default:
		return nil
	}
	return nil
}

func NewRestriction() Factory {
	return &Restriction{}
}


