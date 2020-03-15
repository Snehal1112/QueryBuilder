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

// CreateQuery struct contains the table and fields info to build the create query.
type CreateQuery struct {
	db     *Database
	table  string
	fields []field
	primaryKey string
	foreignKey string

	foreignKeyConstrain *foreignKeyConstrain
}



// Constructor for the CreateQuery.
func NewCreateQuery(db *Database, table string) *CreateQuery {
	return &CreateQuery{db: db, table: table}
}

// Field function used to set the field for create query.
func (c *CreateQuery) Field(name string, fieldType int, length interface{}, constrain []int) *CreateQuery {
	var constrains []string
	for _, v := range constrain {
		constrains = append(constrains, GetConstrain(v))
	}

	fieldDataType := GetDataType(fieldType)
	if IsSupportLength(fieldType) {
		fieldDataType +=  "("+cast.ToString(length)+")"
	}

	c.fields = append(c.fields, field{
		fieldName: name,
		fieldType: fieldDataType,
		constrain: strings.Join(constrains, " "),
	})
	return c
}

// SetPrimaryKey function used to set the PK to multiple columns.
func (c *CreateQuery) SetPrimaryKey(fields []string) *CreateQuery {
	columns := strings.Join(fields, ", ")
	c.primaryKey = fmt.Sprintf("%s (%s)", GetConstrain(PK), columns)
	return c
}

type foreignKeyConstrain struct {
	constrain string
	foreignKey string
	fkTable string
}

func (c *CreateQuery) newForeignKeyConstrain(constrain, foreignKey , fkTable string) *CreateQuery {
	c.foreignKeyConstrain = &foreignKeyConstrain{
		constrain: constrain,
		foreignKey: foreignKey,
		fkTable: fkTable,
	}
	return c
}

func (f *foreignKeyConstrain) onUpdate(referenceOpt int) string {
	return fmt.Sprintf(" ON UPDATE %s", GetReferenceOpt(referenceOpt))
}

func (f *foreignKeyConstrain) onDelete(referenceOpt int) string {
	return fmt.Sprintf(" ON DELETE %s", GetReferenceOpt(referenceOpt))
}

// SetForeignKey set the foreign key on the table.
func (c *CreateQuery) SetForeignKey(onUpdate, onDelete interface{}) *CreateQuery {
	fk := c.foreignKeyConstrain
	c.foreignKey = fmt.Sprintf(", CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s)", fk.constrain, fk.foreignKey, fk.fkTable, fk.foreignKey)
	if onUpdate != nil {
		c.foreignKey += fk.onUpdate(onUpdate.(int))
	}

	if onDelete != nil {
		c.foreignKey += fk.onDelete(onDelete.(int))
	}
	return c
}

func (c *CreateQuery) prepareQuery() string {
	var fields []string
	for _, v := range c.fields {
		fields = append(fields, strings.Trim(v.fieldName+" "+v.fieldType+" "+v.constrain, " "))
	}
	columns := strings.Join(fields, ", ")
	if len(c.primaryKey) != 0 {
		columns += c.primaryKey
	}
	if len(c.foreignKey) != 0 {
		columns += c.foreignKey
	}
	return fmt.Sprintf("CREATE Table IF NOT EXISTS %s ( %s )", c.table, columns)
}

// Execute function execute the create query.
func (c *CreateQuery) Execute() (sql.Result, error) {
	return c.db.Exec(DatabaseQuery, c.prepareQuery(), nil)
}
