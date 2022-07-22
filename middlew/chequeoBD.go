package middlew

import (
	"net/http"
	"github.com/jjvos19/twittor/bd"
)

func ChequeoBD(next http.HandleFunc) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
