package alter

import (
	"bytes"
	"database/sql"
	"html/template"

	"github.com/Snehal1112/QueryBuilder/query"
	"github.com/sirupsen/logrus"
)

// TODO: Use nested templating
const addColTpl = `ADD {{.Name}} {{.FieldType}} {{.Constrain}} {{if .InsertAt.Insert}}{{.InsertAt.Position}}{{if eq .InsertAt.Position "AFTER"}} {{.InsertAt.ExistingColumn}}{{end}}{{end}}`

type AddColumn struct {
	table     *Table
	tableName string
	columns   []column
}

func NewAddColumn(table *Table) AddNewColumn {
	return &AddColumn{tableName: table.name, table: table}
}

func (a *AddColumn) Column(name string, fieldType *query.DataType, constrains *query.Constrain, options ...interface{}) *AddColumn {
	a.columns = append(a.columns, column{
		Name:      name,
		FieldType: fieldType.AsString(),
		Constrain: constrains.AsString(),
	})
	return a
}

func (a *AddColumn) InsertAt(insertAfter bool, existingColumn string) *AddColumn {
	field := &a.columns[len(a.columns)-1]

	var insertAt = "FIRST"
	if insertAfter {
		insertAt = "AFTER"
		field.InsertAt.ExistingColumn = existingColumn
	}

	field.InsertAt.Insert = true
	field.InsertAt.Position = insertAt
	return a
}

func (r *AddColumn) PrepareQuery() string {
	return r.prepareQuery()
}

func (a *AddColumn) prepareQuery() string {
	tpl := template.Must(template.New("Add Columns").Funcs(template.FuncMap{
		"handler": a.table.queryTranspiler,
	}).Parse(queryTpl))

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, map[string]interface{}{
		"table":   a.tableName,
		"columns": a.columns,
		"query": addColTpl,
	}); err != nil {
		logrus.Error(err)
	}

	return buf.String()
}

func (a *AddColumn) Execute() (sql.Result, error) {
	return a.table.Execute()
}
