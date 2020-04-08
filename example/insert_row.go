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

	rows := query.NewRows()
	field1 := query.NewRow()
	field1.SetField("categoryName", "categoryOne")
	rows.SetRow(field1)

	field2 := query.NewRow()
	field2.SetField("categoryName", "categoryTwo")
	rows.SetRow(field2)

	result, err := insert.Into("categories").Rows(rows).Execute()

	if err != nil {
		log.Fatal(err)
	}

	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted new categories into categories table.")
}

func InsertSingleCustomer(builder builder.SQL) {
	table := builder.NewDML().Insert().Into("customers")

	row1 := query.NewRow()
	row1.SetField("name", "snehal")
	row1.SetField("age", 22)
	row1.SetField("address", "V.v.nagar")
	row1.SetField("salary", 5000)

	table.Row(row1).Execute()
}

func InsertCustomers(builder builder.SQL) {
	table := builder.NewDML().Insert().Into("customers")

	rows := query.NewRows()
	row1 := query.NewRow()
	row1.SetField("name", "Ramesh")
	row1.SetField("age", 32)
	row1.SetField("address", "Ahmedabad")
	row1.SetField("salary", 2000)
	rows.SetRow(row1)

	row2 := query.NewRow()
	row2.SetField("name", "Khilan")
	row2.SetField("age", 25)
	row2.SetField("address", "Delhi")
	row2.SetField("salary", 1500)
	rows.SetRow(row2)

	row3 := query.NewRow()
	row3.SetField("name", "kaushik")
	row3.SetField("age", 23)
	row3.SetField("address", "Kota")
	row3.SetField("salary", 2000)
	rows.SetRow(row3)

	row4 := query.NewRow()
	row4.SetField("name", "Chaitali")
	row4.SetField("age", 25)
	row4.SetField("address", "Mumbai")
	row4.SetField("salary", 6500)
	rows.SetRow(row4)

	row5 := query.NewRow()
	row5.SetField("name", "Hardik")
	row5.SetField("age", 27)
	row5.SetField("address", "Bhopal")
	row5.SetField("salary", 8500)
	rows.SetRow(row5)

	row6 := query.NewRow()
	row6.SetField("name", "Komal")
	row6.SetField("age", 22)
	row6.SetField("address", "MP")
	row6.SetField("salary", 4500)
	rows.SetRow(row6)

	row7 := query.NewRow()
	row7.SetField("name", "Muffy")
	row7.SetField("age", 24)
	row7.SetField("address", "Indore")
	row7.SetField("salary", 10000)
	rows.SetRow(row7)

	result, err := table.Rows(rows).Execute()
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

}