package query

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

const (
	FIRST = iota+101
	AFTER
)

type insertAt struct {
	position string
	existingColumn string
}
type newField struct {
	fieldName string
	fieldType string
	constrain string
	insertAt insertAt
}

type ADDTableColumns struct {
	db   *Database
	name string
	fields []newField
}

func NewADDTableColumns(db *Database, name string) *ADDTableColumns {
	return &ADDTableColumns{db: db, name: name}
}

// Field function used to set the field for create query.
func (c *ADDTableColumns) Field(name string, fieldType int, length interface{}, constrains []int) *ADDTableColumns {
	var fieldConstrains []string
	for _, v := range constrains {
		fieldConstrains = append(fieldConstrains, Get(v))
	}

	fieldDataType := GetDataType(fieldType)
	if IsSupportLength(fieldType) {
		fieldDataType +=  "("+cast.ToString(length)+")"
	}

	field := newField{
		fieldName: name,
		fieldType: fieldDataType,
		constrain: strings.Join(fieldConstrains, " "),
	}
	c.fields = append(c.fields, field)
	return c
}

func (c *ADDTableColumns) InsertAt(position int, existingColumn string) *ADDTableColumns {
	field := c.fields[len(c.fields) - 1]

	field.insertAt.position = getInsertPosition(position)
	field.insertAt.existingColumn = existingColumn
	return c
}

func getInsertPosition(position int) string {
	if position == FIRST {
		return "FIRST"
	}
	return "AFTER"
}

func (d *ADDTableColumns) Execute() (sql.Result, error) {
	return d.db.Exec(DatabaseQuery, d.prepareQuery(), nil)
}

func (d *ADDTableColumns) prepareQuery() string {
	var fields []string
	for _, v := range d.fields {
		fields = append(fields, strings.Trim(fmt.Sprintf("%s %s %s",v.fieldName,v.fieldType,v.constrain), " "))
	}
	return fmt.Sprintf("ALTER TABLE %s %s;", d.name, strings.Join(fields, ", "))
}


