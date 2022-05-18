package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/DavidAbellan/effd-backend-GOlang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*SearchProfile encuentra un perfil en la BD*/
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := mongoCn.Database("efdd_database")
	col := db.Collection("user")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("User not found " + err.Error())
		return profile, err
	}
	return profile, nil
}
