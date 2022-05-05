package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/effd-backend-GOlang/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*UserExists recibe un email y devuelve el usuario si existiera*/
func UserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := mongoCn.Database("efdd_database")
	col := db.Collection("user")

	condition := bson.M{"mail": email}
	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	/*formatea el Id en hexadecimal para devolverlo
	como parámetro*/
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
