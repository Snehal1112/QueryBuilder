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

func TestCreateQuery_Field(t *testing.T) {
	setupEnv()
	db := SQLBuilder(os.Getenv("DRIVER"))

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

	createCategories := db.Create("categories")
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

	creatProducts := db.Create("products")
	creatProducts.Field("productId", INT, 50, []int{AI, PK})
	creatProducts.Field("productName", VARCHAR, 225, []int{NOTNULL})
	creatProducts.Field("categoryId", INT, 50, []int{})

	result, err = creatProducts.Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}


}