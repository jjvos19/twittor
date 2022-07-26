package jwt

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jjvos19/twittor/models"
)

/*
   GeneroJWT, genera el jwt.
*/
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MasterDelDesarrollo_grupoFacebook")
	payload := jwt.MapClaims{
		"email": t.Email,
		"nombre": t.Nombre,
		"apellidos": t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia": t.Biografia,
		"ubicacion": t.Ubicacion,
		"sitioweb": t.SitioWeb,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	return tokenStr, err
}
