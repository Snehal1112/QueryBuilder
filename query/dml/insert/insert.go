package insert

import (
	"bytes"
	"database/sql"
	"html/template"
	"log"
	"strings"

	"github.com/Snehal1112/QueryBuilder/query"
	"github.com/sirupsen/logrus"
)

var insertTpl = `INSERT INTO {{.table}} ({{.fields}}) VALUES {{.placeholders}};`

type QueryInsert struct {
	db    *sql.DB
	table string
	row   *query.Row
	rows  query.Rows
}

func NewQueryInsert(table string, db *sql.DB ) Service {
	return &QueryInsert{db: db, table: table}
}

func (q *QueryInsert) Row(fields *query.Row) *QueryInsert{
	q.row = fields
	return q
}

func (q *QueryInsert) Rows(fields query.Rows) *QueryInsert {
	q.rows = fields
	return q
}

func (q *QueryInsert) multiInsertQuery() (string, []interface{}) {
	var placeholders []string

	for _, row := range q.rows {
		_ , placeholder, _ := row.Transpile()
		placeholders = append(placeholders, "(" + strings.Join(placeholder, ", ") + ")")
	}

	tpl := template.Must(template.New("Multi Insert").Parse(insertTpl))
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, map[string]interface{}{
		"table": q.table,
		"fields": strings.Join(q.rows[0].GetNames(), ", "),
		"placeholders": strings.Join(placeholders, ", "),
	}); err != nil {
		logrus.Error(err)
	}

	return buf.String(), q.rows.GetValues()
}

func (q *QueryInsert) singleInsertQuery() (string, []interface{}) {
	fields, placeholders, args := q.row.Transpile()

	tpl := template.Must(template.New("Single Insert").Parse(insertTpl))
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, map[string]interface{}{
		"table": q.table,
		"fields": strings.Join(fields, ", "),
		"placeholders": "(" + strings.Join(placeholders, ", ") + ")",
	}); err != nil {
		logrus.Error(err)
	}

	return buf.String(), args
}

func (q *QueryInsert) prepareQuery() (string, []interface{}) {
	if q.rows != nil && q.row == nil {
		return q.multiInsertQuery()
	}
	return q.singleInsertQuery()
}

func (q *QueryInsert) Execute() (sql.Result, error) {
	queryString, values := q.prepareQuery()

	log.Println(queryString, values)
	return q.db.Exec(queryString, values...)
}
