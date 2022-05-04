package middlew

import (
	"net/http"

	"github.com/DavidAbellan/effd-backend-GOlang/bd"
)

/*DBCheck middleware que chequea la disponibilidad de la base de datos, middleware == recibe lo mismo que retorna */
func DBCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Connection Lost!!", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
