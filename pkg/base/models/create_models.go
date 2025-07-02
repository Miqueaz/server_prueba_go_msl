package base_models

import (
	"log"
	"main/connection/db/postgres"
	helpers "main/pkg/base/helpers"
	query_postgres "main/pkg/postgres"

	"github.com/jmoiron/sqlx"
)

// Crear un nuevo modelo y guardarlo
func NewModel[T any](name string, collectionName string) *Model[T] {
	db := sqlx.NewDb(postgres.DB, "postgres")

	model := &Model[T]{
		Name:           name,
		CollectionName: collectionName,
		Structure:      *new(T),
		Find:           query_postgres.NewQueryBuilder[T](db, collectionName),
	}
	helpers.SaveStructure(model, &models)

	if m, ok := helpers.LoadStructure[Model[T]](&models); ok {
		log.Printf("Modelo '%s' creado y almacenado con éxito.\n", m.Name)
		return model
	}

	return nil
}
