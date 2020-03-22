package query

import (
	"github.com/Snehal1112/QueryBuilder/ddl"
	"github.com/Snehal1112/QueryBuilder/dml"
)

// SQLBuilder struct expost the differnt type of sql querys
type SQLBuilder struct{}

// NewSQLBuilder constructor for the SQLBuilder
func NewSQLBuilder() SQL {
	return &SQLBuilder{}
}

// NewDDL function is entry point for the DDL(Data Definition Language)
func (S *SQLBuilder) NewDDL() ddl.Service {
	return ddl.NewBuilder()
}

// NewDML function is entry point for the DML(Data Manipulation Language)
func (S *SQLBuilder) NewDML() dml.Service {
	return dml.NewBuilder()
}
