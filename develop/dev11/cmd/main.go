package main

import "dev11/internal/server"

func main() {
	srv := server.NewServer()
	srv.SetupHandlers()
	srv.Up()
}
