package main

import (
	"os"

	"QueryBuilder/query"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.Out = os.Stdout

	db := query.NewDatabase("mysql", "root:Snehal@197@/oidb")

	result, err := db.Insert("user").Fields(map[string]interface{}{
		"first_name": "Snehal",
		"last_name": "Dangroshiya",
	}).Execute()

	if err != nil {
		log.Warn(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Warn(err)
	}
	log.Println("Result:", id)
}
