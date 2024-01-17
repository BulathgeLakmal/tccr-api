package main

import (
	"database/sql"
	"log"

	"github.com/it21152832/Learning-Backend/api"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/it21152832/Learning-Backend/util"
)

func main() {
	// Load configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Printf("error loading config: %v", err)
		return
	}

	// Establish a database connection
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Printf("cannot connect to db: %v", err)
		return
	}
	defer conn.Close() // Close the database connection when main exits

	// Create a store instance using the database connection
	store := db.NewStore(conn)
	server,err := api.NewServer(config, store)
	if err != nil{
		log.Fatal("cannot connect to db", err)
	}

	// Create a new server instance by passing the configuration and store
	// server, err := api.NewServer(config, store)
	// if err != nil {
	// 	log.Printf("cannot create server: %v", err)
	// 	return
	// }

	// Start the server using the configured address
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Printf("cannot start server: %v", err)
		return
	}
}
