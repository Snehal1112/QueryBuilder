package insert

import (
	"database/sql"

	"github.com/Snehal1112/QueryBuilder/query"
)

type Service interface {
	prepareQuery() (string, []interface{})
	Execute() (sql.Result, error)
	Row(fields *query.Row) *QueryInsert
	Rows(fields *query.Rows) *QueryInsert
	Select(insertFields []string, selectFields []string, restriction interface{}) *QueryInsert
}