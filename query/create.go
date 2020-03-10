package query

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

type field struct {
	fieldName string
	fieldType string
	constrain string
}

type CreateQuery struct {
	db *Database
	table string
	fields []field
}

func (c *CreateQuery)Field(name string, fieldType string, length int, constrain []int) *CreateQuery {

	var constrains []string
	for _, v := range constrain {
		constrains = append(constrains, GetConstrain(v))
	}
	c.fields = append(c.fields, field{
		fieldName:   name,
		fieldType:   fieldType + "("+cast.ToString(length)+")",
		constrain:   strings.Join(constrains, " "),
	})
	return c
}

func (c *CreateQuery) Execute() (sql.Result, error) {
	var fields []string
	for _, v := range c.fields {
		fields = append(fields, strings.Trim(v.fieldName + " "+ v.fieldType + " "+ v.constrain , " "))
	}
	query := fmt.Sprintf("CREATE Table IF NOT EXISTS %s ( %s )", c.table, strings.Join(fields, ", "))
	return c.db.Exec(DatabaseQuery, query, nil)
}