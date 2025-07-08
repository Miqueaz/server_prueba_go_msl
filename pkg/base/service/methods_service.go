package base_service

import (
	"context"
	"fmt"
	"strconv"
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
	result, err := s.Model.Insert(context.Background(), data)
	return result, err
}

// Update: Actualizar datos en la base de datos
func (s *Service[T]) Update(idStr string, data T) (T, error) {
	id, err := strconv.Atoi(idStr)
	fmt.Printf("Actualizando datos de la colección '%s' con el filtro: %v y los datos: %v\n", id, data)
	// return db.UpdateDocument(filter, data, b.Model.CollectionName) // Implementación real de actualización
	data, err = s.Model.UpdateByID(context.Background(), id, data)
	return data, err

}

// Delete: Eliminar datos de la base de datos
func (s *Service[T]) Delete(idStr string) error {
	id, err := strconv.Atoi(idStr)
	fmt.Printf("Eliminando datos de la colección '%s' con el filtro: %v\n", id)
	// return db.DeleteDocument(filter, b.Model.CollectionName) // Implementación real de eliminación
	_, err = s.Model.DeleteByID(context.Background(), id)
	return err

}

func (s *Service[T]) ReadOne(idStr string) (T, error) {
	// Implementation for reading user data
	id, err := strconv.Atoi(idStr)
	data, err := s.Model.Find.Where("ID", "=", id).Exec(context.Background())
	return data[0], err
}
