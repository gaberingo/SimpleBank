package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gaberingo/SimpleBank/util"
	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config")
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
