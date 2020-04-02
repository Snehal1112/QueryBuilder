package drop

import (
	"bytes"
	"database/sql"
	"html/template"
	"strings"

	"github.com/sirupsen/logrus"
)

const query = `DROP{{if .temporary}} TEMPORARY{{end}} TABLE IF EXISTS {{join  ", " .name}};`

type Table struct {
	names     []string
	db        *sql.DB
	temporary bool
}

type Warnings struct {
	level   string `json:"level"`
	code    int    `json:"code"`
	message string `json:"message"`
}

func NewWarnings() *Warnings {
	return &Warnings{}
}

func NewTable(names []string, db *sql.DB) *Table {
	return &Table{names: names, db: db}
}

func (t *Table) Warning() *Warnings {
	// TODO: get the warning using following sql query
	//  SHOW WARNINGS;
	return NewWarnings()
}

func (w *Warnings) GetErrorCode() int {
	return w.code
}

func (w *Warnings) GetErrorMessage() string {
	return w.message
}

func (t *Table) Temporary(setTemp bool) *Table {
	t.temporary = setTemp
	return t
}

func join(spe string, elem []string) string {
	return strings.Join(elem, spe)
}

func (t *Table) prepareQuery() string {
	// TODO: Improve templating logic
	tpl := template.New("").Funcs(template.FuncMap{"join": join})
	tpl = template.Must(tpl.Parse(query))
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, map[string]interface{}{
		"temporary": t.temporary,
		"name":      t.names,
	}); err != nil {
		logrus.Error("Error in transpile the drop query")
	}
	return buf.String()
}

func (t *Table) Execute() (sql.Result, error) {
	stmt, err := t.db.Prepare(t.prepareQuery())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec()
}
