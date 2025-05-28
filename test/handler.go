package base

import (
	"fmt"
	"server/core/connection/db"
	"server/core/hooks"
)

type Model struct {
	Name           string      `json:"name"`
	CollectionName string      `json:"collectionName"`
	Structure      interface{} `json:"structure"`
}

// BaseModelController maneja operaciones CRUD
type BaseModelController struct {
	hooks.Hookable
	hooks.Cleaners
	Model
}

// Methods define los métodos CRUD
type Methods interface {
	Read(filter map[string]any, config map[string]int) ([]map[string]any, error)
	Insert(data map[string]interface{}) error
	Update(filter map[string]interface{}, data map[string]interface{}) error
	Delete(filter map[string]interface{}) error
}

// Métodos CRUD implementados por BaseModelController

// Método Read con soporte de hooks
func (s *BaseModelController) Read(filter map[string]any, config map[string]int) ([]map[string]any, error) {

	// Ejecutar hooks antes del Read
	if err := s.ExecuteHooks(s.BeforeRead, filter); err != nil {
		return nil, err
	}

	// Simula llamada a la base de datos
	data, err := s.FindClenear(db.FindDocuments(
		filter,
		s.Model.CollectionName,
		int64(config["page"]),
		int64(config["pageSize"]),
	)) // Simula db.FindDocuments

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
