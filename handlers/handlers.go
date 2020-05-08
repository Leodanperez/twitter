package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Leodanperez/twitter/middlew"
	"github.com/Leodanperez/twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manajadores seteo mi puerto, el Handler y pongo a escuchar al Servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
