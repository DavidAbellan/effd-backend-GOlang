package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DavidAbellan/effd-backend-GOlang/bd"
	"github.com/DavidAbellan/effd-backend-GOlang/models"
)

/*ModifyRegister modifica el perfil de usuario */
func ModifyRegister(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid Data "+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModifyUser(t, IdUsuario)
	if err != nil {
		http.Error(w, "Update User error!. Try again "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Cannot update the register "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
