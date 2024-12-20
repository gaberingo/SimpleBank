package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/gaberingo/SimpleBank/api"
	db "github.com/gaberingo/SimpleBank/db/sqlc"
	"github.com/gaberingo/SimpleBank/util"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server :", err)
	}
}
