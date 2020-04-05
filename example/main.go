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

	// fields := query.NewField()
	// fields.SetField("name","snehal")
	// fields.SetField("last_name","dangroshiya")
	// fields.SetField("contact", 9033397218)
	// builders.NewDML().Insert().Into("categories").Row(fields).Execute()

	InsertRow(builders)
	InsertRows(builders)
	//CreateTable(builders)
	//CreateDatabase(builders)
	//DropTable(builders)

	// ddlQuery := builders.NewDDL()
	// result, err := ddlQuery.Alter().Table("categories").Rename().Column("categoryName", "categoryName", query.NewDataType(query.VARCHAR, 200), query.NewConstrain(nil)).Execute()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = result.LastInsertId()
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
