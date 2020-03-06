package query

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(driver string) *Database {
	db, err := sql.Open(driver, dataSourceName())

	if err != nil {
		log.Println(err)
	}
	return &Database{db}
}

func dataSourceName() string {
	if err := godotenv.Load(); err != nil {
		log.Error(err)
	}

	return  fmt.Sprintf("%s:%s@/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
}

func (d *Database)Insert(table string) *InsertQuery {
	return &InsertQuery{db: d, table: table}
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error){
	return d.DB.Exec(query, args...)
}