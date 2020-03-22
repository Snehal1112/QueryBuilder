package drop

import "log"

type Database struct {}

func NewDatabase() *Database {
	return &Database{}
}

func (cd *Database)PrepareQuery() *Database {
	log.Println("PrepareQuery CreateDatabase")
	return cd
}

func (cd *Database)Execute() *Database {
	log.Println("Execute CreateDatabase")
	return cd
}