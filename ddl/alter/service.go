package alter

import "database/sql"

type TableService interface {
	prepareQuery() string
	Execute() (sql.Result, error)
}