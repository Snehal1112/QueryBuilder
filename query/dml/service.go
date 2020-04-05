package dml

import "github.com/Snehal1112/QueryBuilder/query/dml/insert"

type Service interface {
	Insert() InsertService
	Update()
	Delete()
}

type InsertService interface {
	Into(name string) insert.Service
}