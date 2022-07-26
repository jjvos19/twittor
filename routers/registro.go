package routers

import (
	"encoding/json"
	"net/http"
	"github.com/jjvos19/twittor/bd"
	"github.com/jjvos19/twittor/models"
)

/*
   Registro es la funcion para crear en la BD el registro del usuario.
*/
func Registro(w http.ResponseWriter, r * http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		MostrarMensajeError(w, "Error en los datos recibidos", err.Error(), 400)
		//http.Error(w, "Error en los datos recibidos " + err.Error(), 400)
		//return
	}

	if len(t.Email) == 0 {
		MostrarMensajeError(w, "El email del usuario es requerido", "", 400)
		//http.Error(w, "El email del usuario es requerido", 400)
		//return
	}
	if len(t.Password) < 6 {
		MostrarMensajeError(w, "Debe especificar una contrasenha de almenos 6 caracteres", "", 400)
		//http.Error(w, "Debe especificar una contrasenha de almenos 6 caracteres", 400)
		//return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		MostrarMensajeError(w, "Ya existe un usuario registrado con ese email", "", 400)
		//http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		//return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		MostrarMensajeError(w, "Ocurrio un error al intentar realizar el registro de usuario ", err.Error(), 400)
		//http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario " + err.Error())
		//return
	}

	if !status {
		MostrarMensajeError(w, "No se ha logrado insertar el registro del usuario", "", 400)
		//http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		//return
	}

	w.WriteHeader(http.StatusCreated)
}

/*
   MostrarMensajeError muestra el mensaje de error y termina.
*/
func MostrarMensajeError(w http.ResponseWriter, mensaje string, error string, codigo int) {
	msg := mensaje
	if error != "" {
		msg += " " + error
	}
	http.Error(w, msg, codigo)
	return
}
