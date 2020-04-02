package alter

import (
	"bytes"
	"database/sql"
	"html/template"

	"github.com/Snehal1112/QueryBuilder/query"
	"github.com/sirupsen/logrus"
)

// TODO: Use nested templating
const renameColTpl = `CHANGE COLUMN {{.Name}} {{.NewName}} {{.FieldType}} {{.Constrain}}{{if .InsertAt.Insert}}{{.InsertAt.Position}}{{if eq .InsertAt.Position "AFTER"}} {{.InsertAt.ExistingColumn}}{{end}}{{end}}`

type Rename struct {
	table     *Table
	tableName string
	columns   []column
}

func (r *Rename) Table(name, newName string) *Rename {
	panic("implement me")
}

func NewRename(table *Table) RenameItem {
	return &Rename{table: table, tableName:table.name}
}

func (r *Rename) Column(name, newName string, fieldType *query.DataType, constrains *query.Constrain, options ...interface{}) *Rename {
	r.columns = append(r.columns, column{
		Name:      name,
		NewName:   newName,
		FieldType: fieldType.AsString(),
		Constrain: constrains.AsString(),
	})

	return r
}

func (r *Rename) prepareQuery() string {
	tpl := template.Must(template.New("Rename Columns").Funcs(template.FuncMap{
		"handler": r.table.queryTranspiler,
	}).Parse(queryTpl))

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, map[string]interface{}{
		"table":  r.tableName,
		"columns": r.columns,
		"query": renameColTpl,
	}); err != nil {
		logrus.Error(err)
	}
	return buf.String()
}

func (r *Rename) Execute() (sql.Result, error) {
	return r.table.Execute()
}
