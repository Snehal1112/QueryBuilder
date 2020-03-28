package drop

import (
	"testing"
)

var database = NewDatabase("querybuilder", nil)

func TestNewDatabase(t *testing.T) {
	result := database.prepareQuery()
	var want = "DROP DATABASE IF EXISTS querybuilder;"
	if result != want {
		t.Errorf("NewDatabase.prepareQuery returned '%+v', want '%+v'", result, want)
	}
}