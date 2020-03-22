package alter

import "log"

// Table struct
type Table struct{}

// NewTable constructor for the Table struct.
func NewTable() *Table {
	return &Table{}
}

// Fields function use to manage the fields which is going to alter in table.
func (ct *Table) Fields() *Table {
	log.Println("Alter Fields called")
	return ct
}

// PrepareQuery function
func (ct *Table) PrepareQuery() *Table {
	log.Println("Alter PrepareQuery called")
	return ct
}

// Execute function
func (ct *Table) Execute() *Table {
	log.Println("Alter execute called")
	return ct
}
