package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/jjvos19/twittor/middlew"
	"github.com/jjvos19/twittor/routers"
)

/*
   Manejadores seteo mi puerto, el Handler y pongo a escuchar al server
*/
func Manejadores() {
	router := mux.NewRouter()

	// Adicion de los endpoints
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	Port := os.Getenv("PORT_TWITTOR")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}
