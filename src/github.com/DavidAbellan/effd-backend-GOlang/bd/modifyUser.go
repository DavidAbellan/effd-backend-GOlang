package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/effd-backend-GOlang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModifyUser modifica el user*/
func ModifyUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := mongoCn.Database("efdd_database")
	col := db.Collection("user")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Description) > 0 {
		register["description"] = u.Description
	}
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Email) > 0 {
		register["mail"] = u.Email
	}
	register["birthdate"] = u.BirthDate
	uptdString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, uptdString)

	if err != nil {
		return false, err
	}
	return true, nil
}
