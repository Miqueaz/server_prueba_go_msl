package base_service

import (
	"context"
	"fmt"
)

// Método Read con soporte de hooks
func (s *Service[T]) Read(filter map[string]any, config map[string]int) ([]T, error) {

	// Ejecutar hooks antes del Read
	if err := s.ExecuteHooks(s.BeforeRead, filter); err != nil {
		return nil, err
	}

	data, err := s.Model.Find.Exec(context.Background())

	// Ejecutar hooks después del Read
	if err == nil {
		print("NO hay data")
		_ = s.ExecuteHooks(s.AfterRead, filter) // Ignoramos errores de hooks posteriores
	}

	return data, err
}

// Insert: Insertar datos en la base de datos
func (s *Service[T]) Insert(data T) (T, error) {
	fmt.Printf("Insertando datos en la colección '%s': %v\n", data)
	// return db.InsertDocument(data, b.Model.CollectionName) // Implementación real de inserción
	return data, nil
}

// Update: Actualizar datos en la base de datos
func (s *Service[T]) Update(filter map[string]interface{}, data T) error {
	fmt.Printf("Actualizando datos de la colección '%s' con el filtro: %v y los datos: %v\n", filter, data)
	// return db.UpdateDocument(filter, data, b.Model.CollectionName) // Implementación real de actualización
	return nil

}

// Delete: Eliminar datos de la base de datos
func (s *Service[T]) Delete(filter map[string]interface{}) error {
	fmt.Printf("Eliminando datos de la colección '%s' con el filtro: %v\n", filter)
	// return db.DeleteDocument(filter, b.Model.CollectionName) // Implementación real de eliminación
	return nil

}
