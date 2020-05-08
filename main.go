package main

import (
	"log"

	"github.com/Leodanperez/twitter/bd"
	"github.com/Leodanperez/twitter/handlers"
)

func main()  {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return 
	}
	handlers.Manejadores()
}