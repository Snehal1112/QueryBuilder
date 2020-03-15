package query

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Snehal1112/QueryBuilder/constrain"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// Database struct contains the sql DB instance.
type Database struct {
	DB *sql.DB
	isDBSelected bool
}

// SQLBuilder function used create the connect with database.
func SQLBuilder(driver string) *Database {
	db, err := sql.Open(driver, dataSourceName(driver))
	if err != nil {
		log.Println("Error in connection", err)
	}

	if err = db.Ping(); err != nil {
		log.Println("Error is ping :", err)
	}

	database := &Database{DB: db}
	if name := database.GetSelectedDB(); len(name) != 0 {
		database.isDBSelected = true
	}

	return database
}

func dataSourceName(driver string) string {
	// TODO: validate the env and error handling.
	if len(driver) != 0 && driver == "sqlite3" {
		return os.Getenv("SQLITE_DB")
	}

	connectionStr := fmt.Sprintf("%s:%s@/", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	database := os.Getenv("DB_DATABASE")
	if len(database) != 0 {
		connectionStr += fmt.Sprintf("%s", os.Getenv("DB_DATABASE"))
	}
	return connectionStr
}

// SelectDB function used to select the database.
func (d *Database) SelectDB(name string) error{
	_, err := d.DB.Exec("USE "+name)
	if err == nil {
		d.isDBSelected = true
	}
	return err
}

// GetSelectedDB function used to get the selected database.
func (d *Database) GetSelectedDB() string {
	var name string
	d.DB.QueryRow("SELECT DATABASE()").Scan(&name)
	return name
}

// CreateDB function create the database.
func (d *Database) CreateDB(name string) *CreateDatabase {
	return NewCreateDatabase(d, name)
}

// Insert function used to perform the Insert query.
func (d *Database) Insert(tableName string) *InsertQuery {
	return NewInsertQuery(d, tableName)
}

// Create function used to perform the create query.
func (d *Database) CreateTable(name string) *CreateTable {
	return NewCreateQuery(d, name)
}

// Exec function execute the query.
func (d *Database) Exec(queryType int, query string, args ...interface{}) (sql.Result, error) {
	if queryType == constrain.DatabaseQuery {
		stmt, err := d.DB.Prepare(query)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		return stmt.Exec()
	}
	return d.DB.Exec(query, args...)
}

func (d *Database) Close() {
	d.DB.Close()
}