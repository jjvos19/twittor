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
	// Adicion del endpoint para login
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	// Adicion del endpoint para verPerfil
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	// Adicion del endpoint para modificar perfil
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	// Adicionar escritura de tweet
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	// Adicionar leer tweets
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT_TWITTOR")
	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Escuchando en puerto: " + PORT)
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}
