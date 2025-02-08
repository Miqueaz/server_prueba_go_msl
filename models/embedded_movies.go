package models

import (
	Structures "server/core/structures"
	"time"
)

// Movie representa la estructura de una película
type Movie struct {
	ID               string                 `bson:"_id" json:"_id"`                               // ID de la película
	Plot             string                 `bson:"plot" json:"plot"`                             // Sinopsis
	Genres           []string               `bson:"genres" json:"genres"`                         // Géneros
	Runtime          int                    `bson:"runtime" json:"runtime"`                       // Duración en minutos
	Cast             []string               `bson:"cast" json:"cast"`                             // Elenco
	NumMflixComments int                    `bson:"num_mflix_comments" json:"num_mflix_comments"` // Número de comentarios en mflix
	Poster           string                 `bson:"poster" json:"poster"`                         // URL del póster
	Title            string                 `bson:"title" json:"title"`                           // Título de la película
	FullPlot         string                 `bson:"fullplot" json:"fullplot"`                     // Sinopsis completa
	Languages        []string               `bson:"languages" json:"languages"`                   // Idiomas
	Released         *time.Time             `bson:"released" json:"released"`                     // Fecha de estreno
	Directors        []string               `bson:"directors" json:"directors"`                   // Directores
	Writers          []string               `bson:"writers" json:"writers"`                       // Escritores
	Awards           map[string]string      `bson:"awards" json:"awards"`                         // Premios recibidos
	LastUpdated      string                 `bson:"lastupdated" json:"lastupdated"`               // Última fecha de actualización
	Year             int                    `bson:"year" json:"year"`                             // Año de lanzamiento
	IMDB             map[string]interface{} `bson:"imdb" json:"imdb"`                             // Información de IMDB
	Countries        []string               `bson:"countries" json:"countries"`                   // Países de origen
	Type             string                 `bson:"type" json:"type"`                             // Tipo (e.g., película)
	Tomatoes         map[string]interface{} `bson:"tomatoes" json:"tomatoes"`                     // Información de Rotten Tomatoes
	PlotEmbedding    []float64              `bson:"plot_embedding" json:"plot_embedding"`         // Embedding de la sinopsis
}

func init() {
	Structures.NewModel("Movie", "movies", Movie{})
}
