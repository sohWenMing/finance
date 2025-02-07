package databasetests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/sohWenMing/finance/internal/database"
)

var loadedDB *sql.DB

func TestMain(m *testing.M) {
	db, err := database.InitDB("../.env")
	if err != nil {
		log.Fatal(err)
	}

	loadedDB = db
	defer db.Close()

	code := m.Run()
	os.Exit(code)
}

func TestDBPing(t *testing.T) {
	err := loadedDB.Ping()
	if err != nil {
		t.Error(err)
	}
}
