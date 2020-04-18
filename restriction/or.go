package restriction

import (
	"fmt"
	"strings"

	"github.com/Snehal1112/QueryBuilder/query"
)

type Or Condition

func (o *Or) SetGroup(group bool) {
	o.group = group
}

func (o *Or) GetType() int {
	return o.condType
}

func NewOr() ConditionService {
	return &Or{condType: OR}
}

func (o *Or) SetFields(key string, value interface{}, operator Operator) ConditionService {
	o.Fields = append(o.Fields, NewField(query.NewField(key,value), operator))
	return o
}

func (o *Or) ToString() string {
	var andCondition []string

	for _, field := range o.Fields {
		conString := fmt.Sprintf("%s %s %v", field.Field.GetKey(), field.Operator.ToString(),field.Field.GetValue())
		andCondition = append(andCondition, conString)
	}

	condition := strings.Join(andCondition, " OR ")
	if o.group {
		condition = fmt.Sprintf("(%s)", condition)
	}

	return condition
}