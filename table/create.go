package table

import (
	"log"
)

type field struct {
	fieldName string
	fieldType string
	constrain string
}

type CreateTable struct {
	table  string
	fields []field
	primaryKey string
	foreignKey string

}

func NewCreateTable() *CreateTable {
	return &CreateTable{}
}

func (t *CreateTable)Fields()  {
	log.Println("Create table fields")
}

