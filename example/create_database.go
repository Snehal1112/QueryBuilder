package main

import (
	"log"

	builder "github.com/Snehal1112/QueryBuilder"
)

func CreateDatabase(builder builder.SQL){
	ddlQuery := builder.NewDDL()
	createDatabase := ddlQuery.Create().Database("sddd")
	result, err := createDatabase.Execute()

	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("TestDatabase is created")
}
