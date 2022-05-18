package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DavidAbellan/effd-backend-GOlang/bd"
)

/*ShowProfile extrae los valores del user*/
func ShowProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID requested ", http.StatusBadRequest)
		return
	}
	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Cant find the user "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)

}
