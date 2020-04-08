package insert

import (
	"bytes"
	"database/sql"
	"html/template"
	"strings"

	"github.com/Snehal1112/QueryBuilder/query"
	"github.com/sirupsen/logrus"
)

var insertTpl = `INSERT INTO {{.table}} ({{.fields}}) VALUES {{.placeholders}};`
var insertIntoSelectTpl = `INSERT INTO {{.table}} ({{.fields}}) 
SELECT {{.selectField}} FROM {{.selectTable}}`

type QueryInsert struct {
	db    *sql.DB
	table string
	row   *query.Row
	rows  *query.Rows

	insertFromSelect bool
}

func NewQueryInsert(table string, db *sql.DB ) Service {
	return &QueryInsert{db: db, table: table}
}

func (q *QueryInsert) Row(fields *query.Row) *QueryInsert{
	q.row = fields
	return q
}

func (q *QueryInsert) Rows(fields *query.Rows) *QueryInsert {
	q.rows = fields
	return q
}

func (q *QueryInsert) SelectTable(table string, ) *QueryInsert {

	return q
}

func (q *QueryInsert) Select(insertFields []string, selectFields []string, restriction interface{}) *QueryInsert {
	// TODO:-
	//  - table
	// 	- fields
	// 	- selectFields
	// 	- selectTable
	// 	- restriction
	//  Insert into Supplier (supplier_Name, city, country)
	//  select customer_name, city, country from customers;
	//  Insert into ? (?, ?, ?)
	//  select ?, ?, ? from ?;


	return q
}

func (q *QueryInsert) getRows(rows *query.Rows) query.Rows{
	return *rows
}

func (q *QueryInsert) getQuery(qType string, params map[string]interface{}) string {
	tpl := template.Must(template.New(qType).Parse(insertTpl))
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, params); err != nil {
		logrus.Error(err)
	}

	return buf.String()
}

func (q *QueryInsert) multiInsertQuery() (string, []interface{}) {
	var placeholders []string
	for _, row := range q.getRows(q.rows) {
		_ , placeholder, _ := row.Transpile()
		placeholders = append(placeholders, "(" + strings.Join(placeholder, ", ") + ")")
	}

	return q.getQuery("Multi Insert", map[string]interface{}{
		"table": q.table,
		"fields": strings.Join(q.getRows(q.rows)[0].GetNames(), ", "),
		"placeholders": strings.Join(placeholders, ", "),
	}), q.rows.GetValues()
}

func (q *QueryInsert) singleInsertQuery() (string, []interface{}) {
	fields, placeholders, args := q.row.Transpile()
	return q.getQuery("Single Insert", map[string]interface{}{
		"table": q.table,
		"fields": strings.Join(fields, ", "),
		"placeholders": "(" + strings.Join(placeholders, ", ") + ")",
	}), args
}

func (q *QueryInsert) prepareQuery() (string, []interface{}) {
	if q.rows != nil && q.row == nil {
		return q.multiInsertQuery()
	}
	return q.singleInsertQuery()
}

func (q *QueryInsert) Execute() (sql.Result, error) {
	queryString, values := q.prepareQuery()
	stmt, err := q.db.Prepare(queryString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(values...)
}
