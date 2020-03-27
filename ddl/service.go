package ddl

import "github.com/Snehal1112/QueryBuilder/ddl/create"

type Service interface {
	Create() CreateService
	Drop() *DropQuery
	Alter() *AlterQuery
}

type CreateService interface {
	Table(name string) create.TableService
	Database() *create.Database
}

