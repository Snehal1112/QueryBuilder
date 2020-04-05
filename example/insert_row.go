package main

import (
	"log"

	builder "github.com/Snehal1112/QueryBuilder"
	"github.com/Snehal1112/QueryBuilder/query"
)

func InsertRow(builder builder.SQL){
	ddlQuery := builder.NewDML()
	insert := ddlQuery.Insert()

	field1 := query.NewRow()
	field1.SetField("categoryName", "snehalllll")
	result, err :=insert.Into("categories").Row(field1).Execute()

	if err != nil {
		log.Fatal(err)
	}
	_ , err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted new category.")
}


func InsertRows(builder builder.SQL){
	ddlQuery := builder.NewDML()
	insert := ddlQuery.Insert()

	fi := query.NewRows()
	field1 := query.NewRow()
	field1.SetField("categoryName", "categoryOne")
	fi.SetField(field1)

	field2 := query.NewRow()
	field2.SetField("categoryName", "categoryTwo")
	fi.SetField(field2)

	result, err := insert.Into("categories").Rows(*fi).Execute()

	if err != nil {
		log.Fatal(err)
	}

	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted new categories into categories table.")
}
