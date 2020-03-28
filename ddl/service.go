package ddl

import (
	"github.com/Snehal1112/QueryBuilder/ddl/alter"
	"github.com/Snehal1112/QueryBuilder/ddl/create"
	"github.com/Snehal1112/QueryBuilder/ddl/drop"
)

type Service interface {
	Create() CreateService
	Drop() DropService
	Alter() AlterService
}

type CreateService interface {
	Table(name string) create.TableService
	Database(name string) create.DatabaseService
}

type DropService interface {
	Table(name []string) drop.TableService
	Database(name string) drop.DatabaseService
}

type AlterService interface {
	Table(name string) alter.TableService
}
