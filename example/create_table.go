package main

import (
	"log"

	builder "github.com/Snehal1112/QueryBuilder"
	"github.com/Snehal1112/QueryBuilder/query"
)

func CreateTable(builder builder.SQL) {
	// Create table.
	ddlQuery := builder.NewDDL()
	categoriesTable := ddlQuery.Create().Table("categories")

	categoriesTable.Field("categoryId", query.NewDataType(query.INT, 50), query.NewConstrain([]int{query.NOTNULL, query.AI, query.PK}))
	categoriesTable.Field("categoryName", query.NewDataType(query.INT, 50), query.NewConstrain(nil))
	result, err := categoriesTable.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Categories table created ")

	// Associated products table with categories
	productTable := ddlQuery.Create().Table("products")
	productTable.Field("productId", query.NewDataType(query.INT, 50), query.NewConstrain([]int{query.AI, query.PK}))
	productTable.Field("productName", query.NewDataType(query.VARCHAR, 225), query.NewConstrain([]int{query.NOTNULL}))
	productTable.Field("categoryId", query.NewDataType(query.INT, 50), query.NewConstrain(nil))
	productTable.NewForeignKeyConstrain("fk_category", "categoryId", "categories")
	productTable.SetForeignKey(query.Cascade, query.Cascade)
	result, err = categoriesTable.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Products table created ")
}