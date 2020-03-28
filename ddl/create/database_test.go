package create

import (
	"reflect"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	createDB := NewDatabase("querybuilder", nil)

	result := createDB.prepareQuery()
	var want = "CREATE DATABASE IF NOTE EXISTS querybuilder;"
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Database.prepareQuery returned %+v, want %+v", result, want)
	}

}