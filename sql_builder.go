package query

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Snehal1112/QueryBuilder/ddl"
	"github.com/Snehal1112/QueryBuilder/dml"
	_ "github.com/go-sql-driver/mysql"
)

// SQLBuilder struct expos the different type of sql query
type SQLBuilder struct{
	DB *sql.DB
	isDBSelected bool
}

// NewSQLBuilder constructor for the SQLBuilder
func NewSQLBuilder(driver string) SQL {
	log.Println(driver)
	sqlBuilder := &SQLBuilder{}
	db, err := sql.Open(driver, sqlBuilder.dataSourceName(driver))
	if err != nil {
		log.Println("Error in connection", err)
	}

	if err = db.Ping(); err != nil {
		log.Println("Error is ping :", err)
	}
	sqlBuilder.DB = db
	if name := sqlBuilder.GetSelectedDB(); len(name) != 0 {
		sqlBuilder.isDBSelected = true
	}
	return sqlBuilder
}

func (s *SQLBuilder) dataSourceName(driver string) string {
	if len(driver) != 0 && driver == "sqlite3" {
		return os.Getenv("SQLITE_DB")
	}

	connectionStr := fmt.Sprintf("%s:%s@/", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	database := os.Getenv("DB_DATABASE")
	if len(database) != 0 {
		connectionStr += fmt.Sprintf("%s", os.Getenv("DB_DATABASE"))
	}
	log.Println(connectionStr)
	return connectionStr
}

// SelectDB function used to select the database.
func (s *SQLBuilder) SelectDB(name string) error{
	_, err := s.DB.Exec("USE "+name)
	if err == nil {
		s.isDBSelected = true
	}
	return err
}

// GetSelectedDB function used to get the selected database.
func (s *SQLBuilder) GetSelectedDB() string {
	var name string
	s.DB.QueryRow("SELECT DATABASE()").Scan(&name)
	return name
}

// NewDDL function is entry point for the DDL(Data Definition Language)
func (s *SQLBuilder) NewDDL() ddl.Service {
	return ddl.NewBuilder(s.DB)
}

// NewDML function is entry point for the DML(Data Manipulation Language)
func (s *SQLBuilder) NewDML() dml.Service {
	return dml.NewBuilder()
}
