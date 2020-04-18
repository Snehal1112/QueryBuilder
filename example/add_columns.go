package main

import (
	"log"

	builder "github.com/Snehal1112/QueryBuilder"
	"github.com/Snehal1112/QueryBuilder/query"
)

func AddColumns(builder builder.SQL) {
	ddlQuery := builder.NewDDL()
	table := ddlQuery.Alter().Table("customers")
	addColumn := table.Add()
	addColumn.Column("age", query.NewDataType(query.INT,50), query.NewConstrain(nil)).InsertAt(true,"name")

	// result, err := addColumn.Execute()
	// Or
	// result, err := table.Execute()
	result, err := table.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err  = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
}
