package drop

import "log"

type Table struct {}

func NewTable() *Table {
	return &Table{}
}

func (ct *Table)Fields() *Table {
	log.Println("Fields called")
	return ct
}

func (ct *Table)PrepareQuery() *Table {
	log.Println("PrepareQuery called")
	return ct
}

func (ct *Table)Execute() *Table {
	log.Println("execute called")
	return ct
}
