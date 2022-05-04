package main

import (
	"log"

	"github.com/DavidAbellan/effd-backend-GOlang/bd"
	"github.com/DavidAbellan/effd-backend-GOlang/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Cannot Connect Database")
		return
	}
	handlers.RoutesHandler()

}
