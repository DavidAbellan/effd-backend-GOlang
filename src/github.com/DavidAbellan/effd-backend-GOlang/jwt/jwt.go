package jwt

import (
	"time"

	"github.com/DavidAbellan/effd-backend-GOlang/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func JWTGenerator(u models.User) (string, error) {
	key := []byte("ElForoDeDavor_UsuarioAutorizado")

	payload := jwt.MapClaims{
		"email":       u.Email,
		"nombre":      u.Name,
		"descripcion": u.Description,
		"edad":        u.Age,
		"_id":         u.ID.Hex(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
