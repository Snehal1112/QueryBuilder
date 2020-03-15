package query

import (
	"database/sql"
)

type Query interface {
	Execute() (sql.Result, error)
	prepareQuery() string
}
