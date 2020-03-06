package main

import (
	"github.com/joho/godotenv"
	"os"

	"QueryBuilder/query"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.Out = os.Stdout

	if err := godotenv.Load(); err != nil {
		log.Error(err)
	}

	driver := os.Getenv("DRIVER")
	if len(driver) == 0 {
		driver = "mysql"
	}

	log.Println(driver)

	db := query.NewDatabase(driver)

	result, err := db.Insert("user").Fields(map[string]interface{}{
		"first_name": "Snehal",
		"last_name": "Dangroshiya",
	}).Execute()

	if err != nil {
		log.Error(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Warn(err)
	}
	log.Println("Result:", id)
}
