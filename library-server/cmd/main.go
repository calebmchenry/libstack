package main

import (
	"libstack/pkgs/server"
)

func main() {
	s := server.New()
	s.Serve()
}
