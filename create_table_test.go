package query

import (
	"log"
	"os"
	"testing"

	"github.com/Snehal1112/QueryBuilder/constrain"
	"github.com/Snehal1112/QueryBuilder/datatype"
)
func setupEnv() {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "Snehal@1977")
//	os.Setenv("DB_DATABASE", "querybuilder")
	os.Setenv("DRIVER", "mysql")
}

func TestCreateQuery_Field(t *testing.T) {
	setupEnv()
	db := SQLBuilder(os.Getenv("DRIVER"))

	defer db.Close()
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

	if db.isDBSelected == false {
		db.SelectDB("querybuilder")
	}
	createCategories := db.CreateTable("categories")
	createCategories.Field("categoryId", datatype.INT,50, []int{constrain.NOTNULL, constrain.AI, constrain.PK})
	createCategories.Field("categoryName", datatype.VARCHAR, 225, []int{})

	result, err := createCategories.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	creatProducts := db.CreateTable("products")
	creatProducts.Field("productId", datatype.INT, 50, []int{constrain.AI, constrain.PK})
	creatProducts.Field("productName", datatype.VARCHAR, 225, []int{constrain.NOTNULL})
	creatProducts.Field("categoryId", datatype.INT, 50, []int{})
	creatProducts.NewForeignKeyConstrain("fk_category", "categoryId", "categories")
	creatProducts.SetForeignKey(constrain.CASCADE, constrain.CASCADE)

	result, err = creatProducts.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}



}