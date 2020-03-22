package ddl

type Service interface {
	Create() *CreateQuery
	Drop() *DropQuery
	Alter() *AlterQuery
}
