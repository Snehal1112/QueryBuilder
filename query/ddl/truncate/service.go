package truncate

import "database/sql"

type Service interface {
	prepareQuery() string
	Execute() (sql.Result, error)
}