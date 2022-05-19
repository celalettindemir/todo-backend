package main

import (
	"backend/config"
	"backend/server"
	"log"
)

func main() {
	svr := server.NewServer()
	err := svr.StartServer(config.C.Port)
	if err != nil {
		log.Fatalln(err)
	}
}
