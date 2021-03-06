package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/DavidAbellan/effd-backend-GOlang/middlew"
	"github.com/DavidAbellan/effd-backend-GOlang/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*routesHandler setea el puerto, Handler y pone a escuchar al servidor*/
func RoutesHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/question", middlew.DBCheck(middlew.TokenValidate(routers.SetQuestion))).Methods("POST")
	router.HandleFunc("/signin", middlew.DBCheck(routers.SignIn)).Methods("POST")
	router.HandleFunc("/login", middlew.DBCheck(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.DBCheck(middlew.TokenValidate(routers.ShowProfile))).Methods("GET")
	router.HandleFunc("/modprofile", middlew.DBCheck(middlew.TokenValidate(routers.ModifyRegister))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
