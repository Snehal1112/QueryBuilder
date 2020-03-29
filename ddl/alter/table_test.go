package alter

import (
	"testing"

	"github.com/Snehal1112/QueryBuilder/constrain"
	"github.com/Snehal1112/QueryBuilder/datatype"
)

func TestTable_Add(t *testing.T) {
	table := NewTable("categories", nil)
	table.Add().Column("categoryId",datatype.INT,200,[]int{constrain.NOTNULL, constrain.AI}).InsertAt(true, "name").Column("categoryId",datatype.INT,100,[]int{constrain.NOTNULL, constrain.AI}).InsertAt(false, "categoryName")
	table.Execute()
}
