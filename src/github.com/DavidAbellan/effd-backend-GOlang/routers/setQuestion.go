package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DavidAbellan/effd-backend-GOlang/bd"
	"github.com/DavidAbellan/effd-backend-GOlang/models"
)

/*setQuestion en routers, trata la question que viene del navegador*/
func SetQuestion(w http.ResponseWriter, r *http.Request) {
	var message models.Question

	err := json.NewDecoder(r.Body).Decode(&message)
	register := models.SetQuestion{
		UserID:  IDUsuario,
		Message: message.Message,
		Date:    time.Now(),
	}
	_, status, err := bd.InsertQuestion(register)

	if err != nil {
		http.Error(w, "Error inserting question, try again "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Could not insert the question "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
