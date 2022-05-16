package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DavidAbellan/effd-backend-GOlang/bd"
	"github.com/DavidAbellan/effd-backend-GOlang/jwt"
	"github.com/DavidAbellan/effd-backend-GOlang/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "User or/and Password not valid "+err.Error(), 400)
		return
	}
	if len(u.Email) == 0 {
		http.Error(w, "Email required!", 400)
		return
	}
	document, exists := bd.TryLogin(u.Email, u.Password)
	if exists == false {
		http.Error(w, "User exists already!", 400)
		return
	}
	jwtKey, err := jwt.JWTGenerator(document)
	if err != nil {
		http.Error(w, "Token error generation "+err.Error(), 400)
		return
	}
	/*Creamos un nuevo modelo para la respuesta login con el token*/
	resp := models.LoginResponse{
		Token: jwtKey,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*Crear cookies*/
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
