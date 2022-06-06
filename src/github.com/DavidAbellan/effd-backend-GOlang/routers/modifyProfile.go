package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DavidAbellan/effd-backend-GOlang/bd"
	"github.com/DavidAbellan/effd-backend-GOlang/models"
)

/*ModifyRegister modifica el perfil de usuario */
func ModifyRegister(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	fmt.Println("T", t)

	if err != nil {
		http.Error(w, "Invalid Data "+err.Error(), 400)
		return
	}
	var status bool
	fmt.Println("IdUSuar", IDUsuario)

	status, err = bd.ModifyUser(t, IDUsuario)

	if err != nil {
		http.Error(w, "Update User error. Try again "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Cannot update the register "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
