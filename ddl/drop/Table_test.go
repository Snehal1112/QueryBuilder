package drop

import (
	"testing"
)

var table = NewTable([]string{"a", "b"}, nil)

func TestNewTable(t *testing.T) {
	result := table.prepareQuery()
	var want = "DROP TABLE IF EXISTS a, b;"
	if result != want {
		t.Errorf("NewTable.prepareQuery returned %+v, want %+v", result, want)
	}
}

func TestTable_Temporary(t *testing.T) {
	result := table.Temporary(true).prepareQuery()
	var want = "DROP TEMPORARY TABLE IF EXISTS a, b;"
	if result != want {
		t.Errorf("NewTable.prepareQuery returned %+v, want %+v", result, want)
	}
}
