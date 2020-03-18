package query

import (
	"log"
	"os"
	"testing"
)

func TestNewDropTable(t *testing.T) {
	setupEnv()

	db := SQLBuilder(os.Getenv("DRIVER"))

	defer db.Close()

	result, err := db.DropTable([]string{"products", "categories"}).Execute()

	if err != nil {
		log.Fatal(err)
	}
	_, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
}