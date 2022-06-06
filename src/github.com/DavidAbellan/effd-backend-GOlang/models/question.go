package models

/*Question es el modelo de la pregunta*/
type Question struct {
	Message string `bson:"message" json:"message"`
}
