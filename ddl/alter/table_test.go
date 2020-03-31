package alter

import (
	"log"
	"testing"

	"github.com/Snehal1112/QueryBuilder/constrain"
	"github.com/Snehal1112/QueryBuilder/datatype"
)

func TestTable_Add(t *testing.T) {
	table := NewTable("categories", nil)
	qt := table.Add().Column("categoryId", datatype.INT, 200, []int{constrain.NOTNULL, constrain.AI}).InsertAt(true, "name").Column("categoryId", datatype.INT, 100, []int{constrain.NOTNULL, constrain.AI}).InsertAt(false, "").prepareQuery()
	log.Println(qt)
	//col:=table.Rename().Column("sd","sdd", datatype.INT,40, []int{constrain.NOTNULL}).prepareQuery()
	//log.Println(col)
}
