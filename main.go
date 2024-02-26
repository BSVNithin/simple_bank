package main

import (
	"context"
	"log"

	"github.com/BSVNithin/simple_bank/api"
	db "github.com/BSVNithin/simple_bank/db/sqlc"
	"github.com/BSVNithin/simple_bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

const serveraddress = "0.0.0.0:8080"

func main() {
	config, err := util.LoadConfig("./util/..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal(err)
	}
	err = server.Start(serveraddress)
}
