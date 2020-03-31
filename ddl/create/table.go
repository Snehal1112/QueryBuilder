package create

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
	constrain  string
	foreignKey string
	fkTable    string
}

type Table struct {
	db         *sql.DB
	table      string
	fields     []field
	primaryKey string
	foreignKey string

	foreignKeyConstrain *foreignKeyConstrain
}

func NewTable(name string, db *sql.DB) TableService {
	return &Table{db: db, table: name}
}

func (t *Table) Field(name string, fieldType int, length interface{}, constrains []int) *Table {
	var fieldConstrains []string
	for _, v := range constrains {
		fieldConstrains = append(fieldConstrains, constrain.Get(v))
	}

	fieldDataType := datatype.Get(fieldType)
	if datatype.IsSupportLength(fieldType) {
		fieldDataType += "(" + cast.ToString(length) + ")"
	}

	t.fields = append(t.fields, field{
		fieldName: name,
		fieldType: fieldDataType,
		constrain: strings.Join(fieldConstrains, " "),
	})
	return t
}

// SetPrimaryKey function used to set the PK to multiple columns.
func (t *Table) SetPrimaryKey(fields []string) *Table {
	t.primaryKey = fmt.Sprintf("%s (%s)", constrain.Get(constrain.PK), strings.Join(fields, ", "))
	return t
}

func (t *Table) NewForeignKeyConstrain(constrain, foreignKey, fkTable string) *Table {
	t.foreignKeyConstrain = &foreignKeyConstrain{
		constrain:  constrain,
		foreignKey: foreignKey,
		fkTable:    fkTable,
	}
	return t
}

func (f *foreignKeyConstrain) onUpdate(referenceOpt int) string {
	return fmt.Sprintf(" ON UPDATE %s", constrain.GetReferenceOpt(referenceOpt))
}

func (f *foreignKeyConstrain) onDelete(referenceOpt int) string {
	return fmt.Sprintf(" ON DELETE %s", constrain.GetReferenceOpt(referenceOpt))
}

// SetForeignKey set the foreign key on the table.
func (t *Table) SetForeignKey(onUpdate, onDelete interface{}) *Table {
	fk := t.foreignKeyConstrain
	t.foreignKey = fmt.Sprintf(", CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s)", fk.constrain, fk.foreignKey, fk.fkTable, fk.foreignKey)
	if onUpdate != nil {
		t.foreignKey += fk.onUpdate(onUpdate.(int))
	}

	if onDelete != nil {
		t.foreignKey += fk.onDelete(onDelete.(int))
	}
	return t
}

func (t *Table) prepareQuery() string {
	var fields []string
	for _, v := range t.fields {
		fields = append(fields, strings.Trim(v.fieldName+" "+v.fieldType+" "+v.constrain, " "))
	}
	columns := strings.Join(fields, ", ")
	if len(t.primaryKey) != 0 {
		columns += t.primaryKey
	}
	if len(t.foreignKey) != 0 {
		columns += t.foreignKey
	}
	return fmt.Sprintf("CREATE Table IF NOT EXISTS %s ( %s );", t.table, columns)
}

func (t *Table) Execute() (sql.Result, error) {
	stmt, err := t.db.Prepare(t.prepareQuery())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec()
}
