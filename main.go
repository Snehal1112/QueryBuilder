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

	db := query.NewDatabase(driver)

	create := db.Create("dfff")
	create.Field("id", "int", 50, []int{query.NOTNULL, query.AI, query.PK})
	create.Field("first_name", "varchar", 30, []int{})
	create.Field("last_name", "varchar", 30, []int{})

	result, err := create.Execute()
	if err != nil {
		log.Error(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Warn(err)
	}
	log.Println("Result:", id)

	// result, err := db.Insert("user").Fields(map[string]interface{}{
	// 	"first_name": "Snehal",
	// 	"last_name": "Dangroshiya",
	// }).Execute()
	//
	// if err != nil {
	// 	log.Error(err)
	// }
	//
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	log.Warn(err)
	// }
	// log.Println("Result:", id)
}
