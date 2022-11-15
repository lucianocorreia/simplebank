package main

import (
	"database/sql"
	"log"

	"github.com/lucianocorreia/simplebank/api"
	db "github.com/lucianocorreia/simplebank/db/sqlc"

	_ "github.com/lib/pq"
)

var (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:4001"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("could not connect database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
