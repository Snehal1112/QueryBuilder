package query

import (
	"log"
	"testing"
)

func TestNewDropTableColumns(t *testing.T) {
	db:=connectDB()
	defer db.Close()

	table:=db.DropTableColumns("one")
	table.Fields([]interface{}{"f_name", "l_name"})
	table.Field("address")

	query := table.prepareQuery()
	log.Println(query)

}