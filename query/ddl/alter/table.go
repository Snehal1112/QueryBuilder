package alter

import (
	"bytes"
	"database/sql"
	"html/template"
	"log"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	RenameTable = iota + 1
	DropColumns
	RenameColumns
	ModifyColumns
	AddColumns
)

const queryTpl = `ALTER TABLE {{.table}} {{handler .columns .query}}`

type insertAt struct {
	Insert         bool
	Position       string
	ExistingColumn string
}

type column struct {
	Name      string
	FieldType string
	Constrain string
	InsertAt  insertAt
	NewName   string
}

// Table struct
type Table struct {
	name string
	db   *sql.DB

	queryType int
	addCol    AddNewColumn
	rename    RenameItem
}

// NewTable constructor for the Table struct.
func NewTable(name string, db *sql.DB) TableService {
	return &Table{name: name, db: db}
}

func (t *Table)validateCurrentQueryType(queryType int) {
	if t.queryType != 0 {
		logrus.WithFields(logrus.Fields{
			"expectedQuery": queryType,
			"currentQuery":  t.queryType,
		}).Fatal("Already different query type is set")
	}
}
func (t *Table) Add() AddNewColumn {
	t.validateCurrentQueryType(AddColumns)
	t.queryType = AddColumns
	t.addCol = NewAddColumn(t)
	return t.addCol
}

func (t *Table) Rename() RenameItem {
	t.validateCurrentQueryType(RenameColumns)
	t.queryType = RenameColumns
	t.rename = NewRename(t)
	return t.rename
}

func (t *Table) queryTranspiler(columns []column, query string) string {
	var col []string
	tpl := template.Must(template.New("column").Parse(query))
	for _, c := range columns {
		buf := &bytes.Buffer{}
		if err := tpl.Execute(buf, c); err != nil {
			logrus.Error(err)
		}
		col = append(col, buf.String())
	}
	return strings.Join(col, ", ") + ";"
}

// PrepareQuery function
func (t *Table) prepareQuery() string {
	switch t.queryType {
	case DropColumns:
		return ""
	case RenameColumns:
		log.Println("Table prepareQuery called")
		return t.rename.prepareQuery()
	default:
		return t.addCol.prepareQuery()
	}
}

// Execute function
func (t *Table) Execute() (sql.Result, error) {
	t.queryType = 0
	stmt, err := t.db.Prepare(t.prepareQuery())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec()
}
