package query

import (
	"log"
	"os"
	"testing"
)
func setupEnv() {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "Snehal@1977")
	os.Setenv("DB_DATABASE", "querybuilder")
	os.Setenv("DRIVER", "mysql")
}

func connectDB() *Database {
	setupEnv()
	db := SQLBuilder(os.Getenv("DRIVER"))
	return db
}

func TestCreateQuery_Field(t *testing.T) {
	db := connectDB()

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
	createCategories.Field("categoryId", INT,50, []int{NOTNULL, AI, PK})
	createCategories.Field("categoryName", VARCHAR, 225, []int{})

	result, err := createCategories.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	creatProducts := db.CreateTable("products")
	creatProducts.Field("productId", INT, 50, []int{AI, PK})
	creatProducts.Field("productName", VARCHAR, 225, []int{NOTNULL})
	creatProducts.Field("categoryId", INT, 50, []int{})
	creatProducts.NewForeignKeyConstrain("fk_category", "categoryId", "categories")
	creatProducts.SetForeignKey(CASCADE, CASCADE)

	result, err = creatProducts.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}



}