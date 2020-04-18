package truncate

import (
	"testing"
)

func TestNewTable(t *testing.T) {
	table := TableTruncate("customers", nil)
	result := table.prepareQuery()
	var want = "TRUNCATE TABLE customers;"
	if result != want {
		t.Errorf("TableTruncate.prepareQuery returned %+v, want %+v", result, want)
	}

}
