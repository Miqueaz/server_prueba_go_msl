package base_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	base_helpers "main/pkg/base/helpers"
	"strconv"
)

// Método Read con soporte de hooks
func (s *Service[T]) Read(filter map[string]any) ([]T, error) {

	// Ejecutar hooks antes del Read
	if err := s.ExecuteHooks(s.BeforeRead, filter); err != nil {
		return nil, err
	}

	filtros := base_helpers.NormalizarFiltros(filter)

	qb := s.Model.Find
	for campo, cond := range filtros {
		op := fmt.Sprintf("%v", cond[0])
		val := cond[1]
		qb = *qb.Where(campo, op, val)
	}

	data, err := qb.Exec(context.Background())

	// Ejecutar hooks después del Read
	if err == nil {
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

func (s *Service[T]) ReadOne(id int) (T, error) {
	// Implementation for reading user data
	data, err := s.Model.Find.Where("ID", "=", id).Exec(context.Background())
	return data[0], err
}

func Sanitizar[S any, E any](data E) (S, error) {
	// Inicializar un valor vacío de tipo S
	var sanitized S

	// Convertir los datos a JSON para sanitización (por ejemplo, eliminando valores peligrosos)
	sanitizedData, err := json.Marshal(data)
	if err != nil {
		return sanitized, errors.New("failed to sanitize data: " + err.Error())
	}

	// En este punto, podrías modificar o verificar los datos en sanitizedData
	// Aquí se puede agregar lógica para procesar los datos antes de deserializarlos.

	// Deserializar los datos sanitizados de vuelta al tipo `S`
	err = json.Unmarshal(sanitizedData, &sanitized)
	if err != nil {
		return sanitized, errors.New("failed to unmarshal sanitized data: " + err.Error())
	}

	// Retornar los datos sanitizados y sin errores
	return sanitized, nil
}
