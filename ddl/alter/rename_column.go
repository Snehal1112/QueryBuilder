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

// TODO: Use nested templating
const renameColTpl = `CHANGE COLUMN {{.Name}} {{.NewName}} {{.FieldType}} {{.Constrain}} {{if .InsertAt.Insert}}{{.InsertAt.Position}}{{if eq .InsertAt.Position "AFTER"}} {{.InsertAt.ExistingColumn}}{{end}}{{end}}`

type renameColumn struct {
	Name      string
	FieldType string
	Constrain string
	InsertAt  insertAt
	NewName   string
}

type Rename struct {
	table     *Table
	tableName string
	columns   []renameColumn
}

func (r *Rename) Table(name, newName string) *Rename {
	panic("implement me")
}

func NewRename(table *Table) RenameItem {
	return &Rename{table: table, tableName:table.name}
}

func (r *Rename) Column(name, newName string, fieldType int, length interface{}, constrains []int, options ...interface{}) *Rename {
	var fieldConstrains []string

	for _, v := range constrains {
		fieldConstrains = append(fieldConstrains, constrain.Get(v))
	}

	fieldDataType := datatype.Get(fieldType)
	if datatype.IsSupportLength(fieldType) {
		fieldDataType += "(" + cast.ToString(length) + ")"
	}

	r.columns = append(r.columns, renameColumn{
		Name:      name,
		FieldType: fieldDataType,
		NewName:   newName,
		Constrain: strings.Join(fieldConstrains, " "),
	})

	return r
}

func rename(columns []renameColumn) string {
	var col []string
	tpl := template.Must(template.New("col").Parse(renameColTpl))
	for _, c := range columns {
		buf := &bytes.Buffer{}
		if err := tpl.Execute(buf, c); err != nil {
			logrus.Error(err)
		}
		col = append(col, buf.String())
	}
	return strings.Join(col, ", ") + ";"
}

func (r *Rename) prepareQuery() string {
	tpl := template.Must(template.New("Add Columns").Funcs(template.FuncMap{
		"handler": rename,
	}).Parse(queryTpl))

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, map[string]interface{}{
		"table":  r.tableName,
		"columns": r.columns,
	}); err != nil {
		logrus.Error(err)
	}

	return buf.String()}

func (r *Rename) Execute() (sql.Result, error) {
	return r.table.Execute()
}
