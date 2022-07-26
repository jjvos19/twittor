package routers

import (
	"errors"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jjvos19/twittor/bd"
	"github.com/jjvos19/twittor/models"
)

/*
   Email valor del email usado en todos los EndPoints.
*/
var Email string
/*
   IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints.
*/
var IDUsuario string

/*
   ProcesoToken proceso token para extraer sus valores.
*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterDelDesarrollo_grupoFacebook")
	claims := &models.Claim {}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de Token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims,
		func(token *jwt.Token)(interface{}, error) {
			return miClave, nil
		})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Invalido")
	}
	return claims, false, string(""), err
}


