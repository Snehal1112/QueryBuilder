package builder

import (
	"log"
	"os"
	"testing"

	"github.com/Snehal1112/QueryBuilder/query"
)

func setupEnv() {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "Snehal@1977")
	os.Setenv("DB_DATABASE", "querybuilder")
	os.Setenv("DRIVER", "mysql")
}

func TestSQLBuilder(t *testing.T) {
	setupEnv()
	builder := NewSQLBuilder("mysql")

	table := builder.NewDDL().Create().Table("categories")
	table.Field("categoryId", query.INT, 50, []int{query.NOTNULL, query.AI, query.PK})
	table.Field("categoryName", query.VARCHAR, 225, []int{})
	table.Execute()

	builder.NewDDL().Alter().Table("sd").Add().Column("hh", query.NewDataType(query.INT, 50), query.NewConstrain([]int{query.AI, query.NOTNULL})).InsertAt(false, "sd").Execute()
	builder.NewDDL().Drop().Table([]string{"sd"}).Temporary(true).Execute()
}

func TestSQLBuilder_NewDDL(t *testing.T) {
	log.Println("sddddd")
}
