package main

import (
	"log"

	builder "github.com/Snehal1112/QueryBuilder"
	"github.com/Snehal1112/QueryBuilder/query"
)

func CreateTable(builder builder.SQL) {
	// Create table.
	ddlQuery := builder.NewDDL()
	create := ddlQuery.Create()

	/**
		CREATE TABLE categories(
		    categoryId INT AUTO_INCREMENT PRIMARY KEY,
		    categoryName VARCHAR(100) NOT NULL
		);
	*/
	categoriesTable := create.Table("categories")
	categoriesTable.Field("categoryId", query.NewDataType(query.INT, 50), query.NewConstrain([]int{query.NOTNULL, query.AI, query.PK}))
	categoriesTable.Field("categoryName", query.NewDataType(query.VARCHAR, 100), query.NewConstrain(nil))
	result, err := categoriesTable.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Categories table created ")

	/**
		CREATE TABLE products(
		    productId INT AUTO_INCREMENT PRIMARY KEY,
		    productName varchar(100) not null,
		    categoryId INT NOT NULL,
		    CONSTRAINT fk_category
		    FOREIGN KEY (categoryId)
		    REFERENCES categories(categoryId)
		        ON UPDATE CASCADE
		        ON DELETE CASCADE
		)
	 */
	// Associated products table with categories
	productTable := create.Table("products")
	productTable.Field("productId", query.NewDataType(query.INT, 50), query.NewConstrain([]int{query.AI, query.PK}))
	productTable.Field("productName", query.NewDataType(query.VARCHAR, 225), query.NewConstrain([]int{query.NOTNULL}))
	productTable.Field("categoryId", query.NewDataType(query.INT, 50), query.NewConstrain(nil))
	productTable.NewForeignKeyConstrain("fk_category", "categoryId", "categories")
	productTable.SetForeignKey(query.Cascade, query.Cascade)
	result, err = productTable.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Products table created ")
}