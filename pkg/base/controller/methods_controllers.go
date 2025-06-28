package base_controller

import (
	"fmt"
	"main/modules/core/handler/client"
	models "main/pkg/base/models"
	"net/http"
)

// Métodos CRUD implementados por BaseModelController

// Método Read con soporte de hooks
func (s *Controller[T]) Read(res http.ResponseWriter, req *http.Request) {
	ctx, err := client.Init(res, req)
	if err != nil {
		ctx.InternalServerError(err)
		return
	}

	data, _ := s.Service.Methods.Read(nil, nil)

	// Convert []map[string]any to []any
	result := make([]any, len(data))
	for i, v := range data {
		result[i] = v
	}

	ctx.Success("Users fetched successfully", result)
}

// Insert: Insertar datos en la base de datos
func (s *Controller[T]) Insert(res http.ResponseWriter, req *http.Request) {
}

// Update: Actualizar datos en la base de datos
func (s *Controller[T]) Update(res http.ResponseWriter, req *http.Request) {
}

// Delete: Eliminar datos de la base de datos
func (s *Controller[T]) Delete(res http.ResponseWriter, req *http.Request) {
}

// GetModel: Obtener el modelo asociado al controlador
func (s *Controller[T]) GetModel() models.Model[T] {
	fmt.Printf("Obteniendo el modelo asociado al controlador: %s\n", s.Service.Model.Name)
	return s.Service.Model
}
