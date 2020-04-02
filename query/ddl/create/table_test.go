package create

import (
	"testing"

	"github.com/Snehal1112/QueryBuilder/query"
)

func TestNewTable(t *testing.T) {
	/*
		CREATE TABLE categories(
		    categoryId INT AUTO_INCREMENT PRIMARY KEY,
		    categoryName VARCHAR(100) NOT NULL
		) ENGINE=INNODB;

		CREATE TABLE products(
		    productId INT AUTO_INCREMENT PRIMARY KEY,
		    productName varchar(100) not null,
		    categoryId INT,
		    CONSTRAINT fk_category
		    FOREIGN KEY (categoryId)
		        REFERENCES categories(categoryId)
		)

		OR

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
	categoriesTable := NewTable("categories", nil)

	categoriesTable.Field("categoryId", query.NewDataType(query.INT, 50), query.NewConstrain([]int{query.NOTNULL, query.AI, query.PK}))
	categoriesTable.Field("categoryName", query.NewDataType(query.VARCHAR, 225), query.NewConstrain([]int{}))

	var want = "CREATE Table IF NOT EXISTS categories ( categoryId INT(50) NOT NULL AUTO_INCREMENT PRIMARY KEY, categoryName VARCHAR(225) );"
	var result = categoriesTable.prepareQuery()

	if result != want {
		t.Errorf("NewTable.prepareQuery returned %+v, want %+v", result, want)
	}

	// Associated products table with categories
	productTable := NewTable("products", nil)
	productTable.Field("productId", query.NewDataType(query.INT, 50), query.NewConstrain([]int{query.AI, query.PK}))
	productTable.Field("productName", query.NewDataType(query.VARCHAR, 225), query.NewConstrain([]int{query.NOTNULL}))
	productTable.Field("categoryId", query.NewDataType(query.INT, 50), query.NewConstrain(nil))
	productTable.NewForeignKeyConstrain("fk_category", "categoryId", "categories")
	productTable.SetForeignKey(query.Cascade, query.Cascade)

	want = "CREATE Table IF NOT EXISTS products ( productId INT(50) AUTO_INCREMENT PRIMARY KEY, productName VARCHAR(225) NOT NULL, categoryId INT(50), CONSTRAINT fk_category FOREIGN KEY (categoryId) REFERENCES categories(categoryId) ON UPDATE CASCADE ON DELETE CASCADE );"
	result = productTable.prepareQuery()

	if result != want {
		t.Errorf("NewTable.prepareQuery returned %+v, want %+v", result, want)
	}
}
