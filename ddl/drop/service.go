package drop

import "database/sql"

type Service interface {
	prepareQuery() string
	Execute() (sql.Result, error)
}
type TableService interface {
	Warning() *Warnings
	Temporary(setTemp bool) *Table
	Service
}

type DatabaseService interface {
	Service
}