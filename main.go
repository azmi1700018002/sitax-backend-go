package main

import (
	"sitax/config/db"
	"sitax/routes"
)

func main() {

	// Server to database
	db.Server()

	// Initalize the router
	routes.SetupRouter()

	// Run the server
	routes.SetupRouter().Run(":3000")
}
