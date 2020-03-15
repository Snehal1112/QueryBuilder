package query

import (
	"database/sql"
	"strings"

	"github.com/Snehal1112/QueryBuilder/constrain"
)

// InsertQuery struct contains the table, fields information
// which is use to build the insert query.
type InsertQuery struct {
	db     *Database
	table  string
	fields map[string]interface{}
}

// Constructor for the InsertQuery.
func NewInsertQuery(db *Database, table string) *InsertQuery {
	return &InsertQuery{db: db, table: table}
}

// Fields function used to set the fields which further used to prepare the insert query.
func (q *InsertQuery) Fields(fields map[string]interface{}) *InsertQuery {
	q.fields = fields
	return q
}

// Execute function prepare the query and execute that query.
func (q *InsertQuery) Execute() (sql.Result, error) {
	var fields []string
	var args []interface{}
	var placeholders []string

	for k, v := range q.fields {
		fields = append(fields, k)
		args = append(args, v)
		placeholders = append(placeholders, "?")
	}
	query := "INSERT INTO `" + q.table + "` (" + strings.Join(fields, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ")"
	return q.db.Exec(constrain.TableQuery, query, args...)
}
