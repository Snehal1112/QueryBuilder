package main

import (
	"fmt"
	"log"

	builder "github.com/Snehal1112/QueryBuilder"
)

func TruncateTable(builder builder.SQL, tableName string)  {
	truncate := builder.NewDDL().Truncate()
	table := truncate.Table(tableName)
	result, err := table.Execute()
	if err != nil {
		log.Fatal(err)
	}
	output, er := result.RowsAffected()

	if er != nil {
		log.Fatal(er)
	}

	fmt.Sprintf("%s table drop successfully. %d",tableName,  output)
}