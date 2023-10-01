package main

import (
	"taskbtpn/database"
	"taskbtpn/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
