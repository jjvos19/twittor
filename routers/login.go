package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/jjvos19/twittor/bd"
	"github.com/jjvos19/twittor/jwt"
	"github.com/jjvos19/twittor/models"
)

/*
   Login, realiza el login.
*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos " + err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario y/o Contraseña invalidos", 400)
		return
	}

	
}
