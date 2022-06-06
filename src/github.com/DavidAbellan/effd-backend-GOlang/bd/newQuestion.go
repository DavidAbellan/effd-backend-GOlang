package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/effd-backend-GOlang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertQuestion inserta una pregunta en la BD*/
func InsertQuestion(t models.SetQuestion) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := mongoCn.Database("efdd_database")
	col := db.Collection("question")

	register := bson.M{
		"userId":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
