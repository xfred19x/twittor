package main

import (
	"log"

	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/handlers"
)

func main() {

	//me conectare a mi BD
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
