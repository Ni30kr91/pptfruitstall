package main

import (
	"pptfruitstall/server"
	"pptfruitstall/server/db"
)

func main() {
	db.DB_connection()
	server.Server()
}
