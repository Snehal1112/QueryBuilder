package query

import (
	"database/sql"
	"strings"
)

type InsertQuery struct {
	db *Database
	table string
	fields map[string]interface{}
}

func (q *InsertQuery) Fields(fields map[string]interface{}) *InsertQuery {
	q.fields = fields
	return q
}

func (q *InsertQuery) Execute() (sql.Result, error ){
	var fields []string
	var args []interface{}
	var placeholders []string

	for k, v :=range q.fields{
		fields = append(fields, k)
		args = append(args,v)
		placeholders = append(placeholders, "?")
	}
	query := "INSERT INTO `"+q.table+"` ("+strings.Join(fields, ", ")+") VALUES ("+strings.Join(placeholders, ", ")+")"

	return q.db.Exec(query, args...)
}