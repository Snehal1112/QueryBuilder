package query

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Snehal1112/QueryBuilder/constrain"
	"github.com/Snehal1112/QueryBuilder/datatype"
	"github.com/spf13/cast"
)

type field struct {
	fieldName string
	fieldType string
	constrain string
}

type foreignKeyConstrain struct {
	constrain string
	foreignKey string
	fkTable string
}

// CreateQuery struct contains the table and fields info to build the create query.
type CreateTable struct {
	db     *Database
	table  string
	fields []field
	primaryKey string
	foreignKey string

	foreignKeyConstrain *foreignKeyConstrain
}

// Constructor for the CreateQuery.
func NewCreateQuery(db *Database, table string) *CreateTable {
	return &CreateTable{db: db, table: table}
}

// Field function used to set the field for create query.
func (c *CreateTable) Field(name string, fieldType int, length interface{}, constrains []int) *CreateTable {
	var fieldConstrains []string
	for _, v := range constrains {
		fieldConstrains = append(fieldConstrains, constrain.Get(v))
	}

	fieldDataType := datatype.GetDataType(fieldType)
	if datatype.IsSupportLength(fieldType) {
		fieldDataType +=  "("+cast.ToString(length)+")"
	}

	c.fields = append(c.fields, field{
		fieldName: name,
		fieldType: fieldDataType,
		constrain: strings.Join(fieldConstrains, " "),
	})
	return c
}

// SetPrimaryKey function used to set the PK to multiple columns.
func (c *CreateTable) SetPrimaryKey(fields []string) *CreateTable {
	c.primaryKey = fmt.Sprintf("%s (%s)", constrain.Get(constrain.PK), strings.Join(fields, ", "))
	return c
}

func (c *CreateTable) NewForeignKeyConstrain(constrain, foreignKey , fkTable string) *CreateTable {
	c.foreignKeyConstrain = &foreignKeyConstrain{
		constrain: constrain,
		foreignKey: foreignKey,
		fkTable: fkTable,
	}
	return c
}

func (f *foreignKeyConstrain) onUpdate(referenceOpt int) string {
	return fmt.Sprintf(" ON UPDATE %s", constrain.GetReferenceOpt(referenceOpt))
}

func (f *foreignKeyConstrain) onDelete(referenceOpt int) string {
	return fmt.Sprintf(" ON DELETE %s", constrain.GetReferenceOpt(referenceOpt))
}

// SetForeignKey set the foreign key on the table.
func (c *CreateTable) SetForeignKey(onUpdate, onDelete interface{}) *CreateTable {
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

func (c *CreateTable) prepareQuery() string {
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
func (c *CreateTable) Execute() (sql.Result, error) {
	return c.db.Exec(constrain.DatabaseQuery, c.prepareQuery(), nil)
}
