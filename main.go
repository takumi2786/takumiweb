package main

import (
	"github.com/takumi2786/takumiweb/db"
	"github.com/takumi2786/takumiweb/server"
)

func main() {
	db.Init()
	defer db.Close()
	server.Init()
}
