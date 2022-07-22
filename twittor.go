package main

import (
	"log"
	"github.com/jjvos19/twittor/handlers"
	"github.com/jjvos19/twittor/bd"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}
	
	handlers.Manejadores()
}
