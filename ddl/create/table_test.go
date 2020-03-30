package create

import (
	"testing"

	"github.com/Snehal1112/QueryBuilder/constrain"
	"github.com/Snehal1112/QueryBuilder/datatype"
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
	categoriesTable.Field("categoryId", datatype.INT, 50, []int{constrain.NOTNULL, constrain.AI, constrain.PK})
	categoriesTable.Field("categoryName", datatype.VARCHAR, 225, []int{})

	var want = "CREATE Table IF NOT EXISTS categories ( categoryId INT(50) NOT NULL AUTO_INCREMENT PRIMARY KEY, categoryName VARCHAR(225) );"
	var result = categoriesTable.prepareQuery()

	if result != want {
		t.Errorf("NewTable.prepareQuery returned %+v, want %+v", result, want)
	}

	// Associated products table with categories
	productTable := NewTable("products", nil)
	productTable.Field("productId", datatype.INT, 50, []int{constrain.AI, constrain.PK})
	productTable.Field("productName", datatype.VARCHAR, 225, []int{constrain.NOTNULL})
	productTable.Field("categoryId", datatype.INT, 50, []int{})
	productTable.NewForeignKeyConstrain("fk_category", "categoryId", "categories")
	productTable.SetForeignKey(constrain.Cascade, constrain.Cascade)

	want = "CREATE Table IF NOT EXISTS products ( productId INT(50) AUTO_INCREMENT PRIMARY KEY, productName VARCHAR(225) NOT NULL, categoryId INT(50), CONSTRAINT fk_category FOREIGN KEY (categoryId) REFERENCES categories(categoryId) ON UPDATE CASCADE ON DELETE CASCADE );"
	result = productTable.prepareQuery()

	if result != want {
		t.Errorf("NewTable.prepareQuery returned %+v, want %+v", result, want)
	}
}
