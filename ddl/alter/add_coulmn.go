package alter

import (
	"bytes"
	"database/sql"
	"html/template"
	"strings"

	"github.com/Snehal1112/QueryBuilder/constrain"
	"github.com/Snehal1112/QueryBuilder/datatype"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

const addColTpl = `ADD {{.Name}} {{.FieldType}} {{.Constrain}} {{if .InsertAt.Insert}}{{.InsertAt.Position}} {{.InsertAt.ExistingColumn}}{{end}}`
const queryTpl = `ALTER TABLE {{.table}} {{addColumn .columns}}`

type insertAt struct {
	Insert bool
	Position string
	ExistingColumn string
}

type column struct {
	Name string
	FieldType string
	Constrain string
	InsertAt insertAt
}

type AddColumn struct {
	table *Table
	tableName string
	columns []column
}

func NewAddColumn(table *Table) AddNewColumn {
	return &AddColumn{tableName: table.name, table:table}
}

func (a *AddColumn) Column(name string, fieldType int, length interface{}, constrains []int, options ...interface{}) *AddColumn {
	var fieldConstrains []string

	for _, v := range constrains {
		fieldConstrains = append(fieldConstrains, constrain.Get(v))
	}

	fieldDataType := datatype.Get(fieldType)
	if datatype.IsSupportLength(fieldType) {
		fieldDataType +=  "("+cast.ToString(length)+")"
	}

	a.columns = append(a.columns, column{
		Name: name,
		FieldType: fieldDataType,
		Constrain: strings.Join(fieldConstrains, " "),
	})
	return a
}

func (a *AddColumn) InsertAt(insertFirst bool, existingColumn string) *AddColumn {
	field := &a.columns[len(a.columns) - 1]

	var insertAt = "AFTER"
	if insertFirst {
		insertAt = "FIRST"
	}

	field.InsertAt.Insert = true
	field.InsertAt.Position = insertAt
	field.InsertAt.ExistingColumn = existingColumn
	return a
}

func addColumn(columns []column) string{
	var col []string
	for _, c := range columns {

		tpl := template.Must(template.New("col").Parse(addColTpl))
		buf := &bytes.Buffer{}
		if err := tpl.Execute(buf, c); err != nil {
			logrus.Error(err)
		}
		col = append(col, buf.String())
	}
	return strings.Join(col, ", ") + ";"
}

func (a *AddColumn) prepareQuery() string {
	tpl := template.Must(template.New("Add Columns").Funcs(template.FuncMap{
		"addColumn" : addColumn,
	}).Parse(queryTpl))

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, map[string]interface{}{
		"table" : a.tableName,
		"columns" : a.columns,
	}); err != nil {
		logrus.Error(err)
	}

	return buf.String()
}

func (a *AddColumn) Execute() (sql.Result, error) {
	return a.table.Execute()
}
