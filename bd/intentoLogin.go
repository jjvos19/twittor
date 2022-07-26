package bd

import (
	"github.com/jjvos19/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*
   IntentoLogin realiza el chequeo del login a la BD
*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err !=  nil {
		return usu, false
	}

	return usu, true
}
