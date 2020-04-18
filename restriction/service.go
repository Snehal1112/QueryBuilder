package restriction

import (
	"sync"

	"github.com/Snehal1112/QueryBuilder/query"
)

type Field struct {
	Field query.FieldService
	Operator Operator
}

func NewField(fields query.FieldService, operator Operator) *Field {
	return &Field{Field: fields, Operator: operator}
}

type Condition struct {
	sync.RWMutex
	Fields []*Field
	condType int
	group bool
}

type ConditionService interface {
	SetFields(key string, value interface{}, operator Operator) ConditionService
	GetType() int
	SetGroup(bool)
	ToString() string
}

type Factory interface {
	Factory(restrictionType int) ConditionService
	AddCondition(fn func()string, operator Operator)
	Transpile() string
}
