package main

import (
	"libstack/pkgs/impl/server"
)

func main() {
	s := server.New()
	s.Serve()
}
