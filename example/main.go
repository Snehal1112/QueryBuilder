package main

import (
	"os"

	builder "github.com/Snehal1112/QueryBuilder"
)

func main()  {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "Snehal@1977")
	os.Setenv("DB_DATABASE", "querybuilder")
	os.Setenv("DRIVER", "mysql")

	builders := builder.NewSQLBuilder("mysql")
	//InsertRow(builders)
	//InsertRows(builders)
	//CreateTable(builders)
	//CreateDatabase(builders)
	//DropTable(builders)

	CreateCustomersTable(builders)
	AddColumns(builders)
	InsertCustomers(builders)
	InsertSingleCustomer(builders)
	TruncateTable(builders, "customers")
}
