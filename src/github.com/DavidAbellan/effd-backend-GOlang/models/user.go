package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name,omitempty"`
	BirthDate   time.Time          `bson:"birth" json:"birth,omitempty"`
	Email       string             `bson:"mail,omitempty" json:"mail"`
	Description string             `bson:"description" json:"description,omitempty"`
	Age         int32              `bson:"age" json:"age,omitempty"`
	Password    string             `bson:"password" json:"password,omitempty"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempty"`
}
