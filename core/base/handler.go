package base

import (
	"fmt"
	"server/core/connection/db"

	"server/core/hooks"

	"go.mongodb.org/mongo-driver/mongo"
)

// BaseModelController maneja operaciones CRUD
type BaseModelController struct {
	Model
	hooks.Hookable
}

// Methods define los métodos CRUD
type Methods interface {
	Read(filter map[string]interface{}, config map[string]int) (*mongo.Cursor, error)
	Insert(data map[string]interface{}) error
	Update(filter map[string]interface{}, data map[string]interface{}) error
	Delete(filter map[string]interface{}) error
}

// Métodos CRUD implementados por BaseModelController

// Método Read con soporte de hooks
func (s *BaseModelController) Read(filter map[string]interface{}, config map[string]int) (*mongo.Cursor, error) {
	fmt.Printf("Reading from collection '%s' with filter: %v\n", s.Model.CollectionName, filter)

	// Ejecutar hooks antes del Read
	if err := s.ExecuteHooks(s.BeforeRead, filter); err != nil {
		return nil, err
	}

	// Simula llamada a la base de datos
	data, err := db.FindDocuments(filter, s.Model.CollectionName, int64(config["page"]), int64(config["pageSize"])) // Simula db.FindDocuments

	// Ejecutar hooks después del Read
	if err == nil {
		_ = s.ExecuteHooks(s.AfterRead, filter) // Ignoramos errores de hooks posteriores
	}

	return data, err
}

// Insert: Insertar datos en la base de datos
func (b *BaseModelController) Insert(data map[string]interface{}) error {
	fmt.Printf("Insertando datos en la colección '%s': %v\n", b.Model.CollectionName, data)
	// return db.InsertDocument(data, b.Model.CollectionName) // Implementación real de inserción
	return nil
}

// Update: Actualizar datos en la base de datos
func (b *BaseModelController) Update(filter map[string]interface{}, data map[string]interface{}) error {
	fmt.Printf("Actualizando datos de la colección '%s' con el filtro: %v y los datos: %v\n", b.Model.CollectionName, filter, data)
	// return db.UpdateDocument(filter, data, b.Model.CollectionName) // Implementación real de actualización
	return nil

}

// Delete: Eliminar datos de la base de datos
func (b *BaseModelController) Delete(filter map[string]interface{}) error {
	fmt.Printf("Eliminando datos de la colección '%s' con el filtro: %v\n", b.Model.CollectionName, filter)
	// return db.DeleteDocument(filter, b.Model.CollectionName) // Implementación real de eliminación
	return nil

}
