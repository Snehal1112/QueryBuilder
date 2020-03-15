package table

import (
	"log"
)


type CreateDatabase struct {
	dbName string
}

func NewCreateDatabase() *CreateDatabase {
	return &CreateDatabase{}
}

func (t *CreateDatabase)Fields()  {
	log.Println("Create database fields")
}

