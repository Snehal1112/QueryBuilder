package alter

import (
	"bytes"
	"database/sql"
	"html/template"

	"github.com/Snehal1112/QueryBuilder/query"
	"github.com/sirupsen/logrus"
)

// TODO: Use nested templating
const reColumnTpl = `CHANGE COLUMN {{.Name}} {{.NewName}} {{.FieldType}} {{.Constrain}}{{if .InsertAt.Insert}}{{.InsertAt.Position}}{{if eq .InsertAt.Position "AFTER"}} {{.InsertAt.ExistingColumn}}{{end}}{{end}}`
const reTableTpl = `ALTER TABLE {{.table}} RENAME TO {{.newName}};`

const (
	ReColumn = iota + 0
	ReTable
)

type Rename struct {
	table     *Table
	columns   []column
	newTableName string
	renameType int
}

func (r *Rename) Table(newName string) *Rename {
	r.renameType = ReTable
	r.newTableName = newName
	return r
}

func NewRename(table *Table) RenameItem {
	return &Rename{table: table}
}

func (r *Rename) Column(name, newName string, fieldType *query.DataType, constrains *query.Constrain, options ...interface{}) *Rename {
	r.columns = append(r.columns, column{
		Name:      name,
		NewName:   newName,
		FieldType: fieldType.AsString(),
		Constrain: constrains.AsString(),
	})
	r.renameType = ReColumn
	return r
}

func (r *Rename) executeTpl(tpl *template.Template, queryData map[string]interface{}) string {
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, queryData); err != nil {
		logrus.Error(err)
	}
	return buf.String()
}

func (r *Rename) prepareQueryToRenameColumn() string {
	tpl := template.Must(template.New("Rename Column").Funcs(template.FuncMap{
		"handler": r.table.queryTranspiler,
	}).Parse(queryTpl))

	return r.executeTpl(tpl, map[string]interface{}{
		"table":  r.table.name,
		"columns": r.columns,
		"query": reColumnTpl,
	})
}

func (r *Rename) prepareQueryToRenameTable() string {
	tpl := template.Must(template.New("Rename Table").Parse(reTableTpl))

	return r.executeTpl(tpl, map[string]interface{}{
		"table":  r.table.name,
		"newName": r.newTableName,
	})
}

func (r *Rename) prepareQuery() string {
	if r.renameType == ReColumn {
		return r.prepareQueryToRenameColumn()
	}
	return r.prepareQueryToRenameTable()
}

func (r *Rename) Execute() (sql.Result, error) {
	return r.table.Execute()
}
