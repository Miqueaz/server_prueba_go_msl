package models

import (
	Structures "server/core/structures"
	"time"
)

// Nombre de la colección en la base de datos
var nombre string = "Comentarios"
var collectionName string = "comments"

// Comment representa un comentario en la base de datos
type commentModel struct {
	ID      string    `bson:"_id" json:"_id"`           // ID del comentario
	Name    string    `bson:"name" json:"name"`         // Nombre del autor
	Email   string    `bson:"email" json:"email"`       // Email del autor
	MovieID string    `bson:"movie_id" json:"movie_id"` // ID de la película asociada
	Text    string    `bson:"text" json:"text"`         // Contenido del comentario
	Date    time.Time `bson:"date" json:"date"`         // Fecha del comentario
}

func init() {
	Structures.NewModel(nombre, collectionName, commentModel{})
}
