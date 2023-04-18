package main

import (
	"database/sql"
	"log"

	"github.com/aryzk29/simplebankcp/api"
	db "github.com/aryzk29/simplebankcp/db/sqlc"
	"github.com/aryzk29/simplebankcp/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server: ", err)
	}

}
