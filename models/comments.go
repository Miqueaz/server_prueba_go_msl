package models

import (
	"time"
)

// Comment representa un comentario en la base de datos
type Comment struct {
	ID      string    `bson:"_id" json:"_id"`           // ID del comentario
	Name    string    `bson:"name" json:"name"`         // Nombre del autor
	Email   string    `bson:"email" json:"email"`       // Email del autor
	MovieID string    `bson:"movie_id" json:"movie_id"` // ID de la película asociada
	Text    string    `bson:"text" json:"text"`         // Contenido del comentario
	Date    time.Time `bson:"date" json:"date"`         // Fecha del comentario
}
