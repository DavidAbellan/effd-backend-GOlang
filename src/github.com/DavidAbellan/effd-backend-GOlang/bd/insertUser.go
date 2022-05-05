package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/effd-backend-GOlang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertUser inserta un usuario en la bd*/
func InsertUser(u models.User) (string, bool, error) {
	/*
	  Se crea el contexto de la bd con un timeOut de 15 seg, desde el contexto anterior
	  defer se ejecuta al final de la funcion y cancela el contexto de 15 seg
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := mongoCn.Database("efdd_database")
	col := db.Collection("user")

	u.Password, _ = EncryptThePassword(u.Password)
	u.Validation = false
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	/*
	   La línea 30 es para tratar el Id que devuelve mongo despues del insert
	   y poder devolverlo en string
	*/
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
