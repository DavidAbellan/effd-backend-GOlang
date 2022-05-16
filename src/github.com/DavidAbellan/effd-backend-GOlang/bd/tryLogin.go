package bd

import (
	"github.com/DavidAbellan/effd-backend-GOlang/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin intenta loguearse en el sistema*/
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := UserExists(email)
	if found == false {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true

}
