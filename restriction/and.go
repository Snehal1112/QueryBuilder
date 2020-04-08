package restriction

import (
	"fmt"
	"strings"

	"github.com/Snehal1112/QueryBuilder/query"
)


type And Condition

func (a *And) SetGroup( group bool) {
	a.group = group
}

func (a *And) GetType() int {
	return a.condType
}

func (a *And) SetFields(key string, value interface{}, operator Operator) ConditionService{
	a.Lock()
	defer a.Unlock()
	a.Fields = append(a.Fields, NewField(query.NewField(key,value), operator))
	return a
}

func (a *And) ToString() string {
	var andCondition []string

	for _, field := range a.Fields {
		conString := fmt.Sprintf("%s %s %v", field.Field.GetKey(), field.Operator.ToString(), field.Field.GetValue())
		andCondition = append(andCondition, conString)
	}

	condition := strings.Join(andCondition, " AND ")
	if a.group {
		condition = fmt.Sprintf("(%s)", condition)
	}

	return condition
}

func NewAnd() ConditionService {
	return &And{condType:AND}
}