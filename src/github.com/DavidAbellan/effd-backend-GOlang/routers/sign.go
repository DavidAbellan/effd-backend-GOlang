package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DavidAbellan/effd-backend-GOlang/bd"
	"github.com/DavidAbellan/effd-backend-GOlang/models"
)

/*SignIn es la función para crear el registro en la bd del user*/
func SignIn(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Data Error "+err.Error(), 400)
		return
	}
	if len(u.Email) == 0 {
		http.Error(w, "Mandatory email "+err.Error(), 400)
		return
	}
	if len(u.Password) < 6 {
		http.Error(w, "Password must be longer than six characters "+err.Error(), 400)
		return
	}

	_, found, _ := bd.UserExists(u.Email)
	if found == true {
		http.Error(w, "User registered already "+err.Error(), 400)
		return
	}

	_, status, err := bd.InsertUser(u)
	if err != nil {
		http.Error(w, "Register Error !! "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Status Error Register Failed "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
