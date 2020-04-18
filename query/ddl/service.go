package ddl

import (
	"github.com/Snehal1112/QueryBuilder/query/ddl/alter"
	"github.com/Snehal1112/QueryBuilder/query/ddl/create"
	"github.com/Snehal1112/QueryBuilder/query/ddl/drop"
	"github.com/Snehal1112/QueryBuilder/query/ddl/truncate"
)

type Service interface {
	Create() CreateService
	Drop() DropService
	Alter() AlterService
	Truncate() TruncateService
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

type TruncateService interface {
	Table(name string) truncate.Service
}
