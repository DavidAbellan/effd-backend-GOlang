package routers

import (
	"errors"
	"strings"

	"github.com/DavidAbellan/effd-backend-GOlang/bd"
	"github.com/DavidAbellan/effd-backend-GOlang/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email y IdUsuario son variables globales para poder usar en todos los endpoints*/
var Email string
var IDUsuario string

/*ProcessToken desencripta el token y comprueba que es valido*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	key := []byte("ElForoDeDavor_UsuarioAutorizado")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token invalid format")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		_, found, _ := bd.UserExists(claims.Email)
		if found {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, found, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err

}
