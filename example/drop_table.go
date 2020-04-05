package main

import (
	"log"

	builder "github.com/Snehal1112/QueryBuilder"
)

func DropTable(builder builder.SQL)  {
	drop := builder.NewDDL().Drop()
	table := drop.Table([]string{"TESTDB"})
	result, err := table.Execute()
	if err != nil {
		log.Fatal(err)
	}

	output, er := result.RowsAffected()

	if er != nil {
		log.Fatal(er)
	}

	log.Println("products table drop successfully.", output)
}