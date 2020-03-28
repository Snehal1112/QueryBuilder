package query

import (
	"log"
	"os"
	"testing"

	"github.com/Snehal1112/QueryBuilder/constrain"
	"github.com/Snehal1112/QueryBuilder/datatype"
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

	table :=builder.NewDDL().Create().Table("categories")
	table.Field("categoryId", datatype.INT,50, []int{constrain.NOTNULL, constrain.AI, constrain.PK})
	table.Field("categoryName", datatype.VARCHAR, 225, []int{})
	table.Execute()

	builder.NewDDL().Drop().Table([]string{"sd"}).Temporary(true).Execute()
}

func TestSQLBuilder_NewDDL(t *testing.T) {
	log.Println("sddddd")
}