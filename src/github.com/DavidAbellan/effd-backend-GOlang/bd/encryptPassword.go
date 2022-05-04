package bd

import "golang.org/x/crypto/bcrypt"

/*EncryptThePassword encripta el password para insertarlo en la bd*/
func EncryptThePassword(pass string) (string, error) {
	cost := 6
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(hash), err
}
