package alter

import (
	"testing"

	"github.com/Snehal1112/QueryBuilder/query"
)


func TestTable_Add(t *testing.T) {
	var table = NewTable("categories", nil)
	var want = "ALTER TABLE categories ADD categoryId INT(200) NOT NULL AUTO_INCREMENT AFTER name, ADD categoryId INT(200) NOT NULL AUTO_INCREMENT FIRST;"
	addCol := table.Add()
	addCol.Column("categoryId", query.NewDataType(query.INT, 200), query.NewConstrain([]int{query.NOTNULL, query.AI})).InsertAt(true, "name")
	addCol.Column("categoryId", query.NewDataType(query.INT, 200), query.NewConstrain([]int{query.NOTNULL, query.AI})).InsertAt(false, "")
	result := addCol.prepareQuery()

	if result != want {
		t.Errorf("NewTable.Add.prepareQuery returned %+v, want %+v", result, want)
	}
}

func TestTable_Rename(t *testing.T) {
	var table = NewTable("categories", nil)
	rename := table.Rename()
	rename.Column("sd","sdd", query.NewDataType(query.INT, 200), query.NewConstrain([]int{query.NOTNULL, query.AI}))
	result := rename.prepareQuery()
	var want = "ALTER TABLE categories CHANGE COLUMN sd sdd INT(200) NOT NULL AUTO_INCREMENT;"

	if result != want {
		t.Errorf("NewTable.Add.prepareQuery returned %+v, want %+v", result, want)
	}
}

func TestRename_Table(t *testing.T) {
	var table = NewTable("categories", nil)
	rename := table.Rename().Table("sddd")
	result := rename.prepareQuery()
	var want = "ALTER TABLE categories RENAME TO sddd;"

	if result != want {
		t.Errorf("NewTable.Add.prepareQuery returned %+v, want %+v", result, want)
	}
}