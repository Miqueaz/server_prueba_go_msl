package Structures

import (
	"fmt"
	"server/core/connection/db"
)

// BaseModelController para manejar operaciones CRUD
type BaseModelController struct {
	Model *Model
}

type Methods interface {
	Read(filter map[string]interface{}) (interface{}, error)
	Insert(data map[string]interface{}) error
	Update(filter map[string]interface{}, data map[string]interface{}) error
	Delete(filter map[string]interface{}) error
}

// Método Read: Leer documentos de la base de datos
func (s *BaseModelController) Read(filter map[string]interface{}) (interface{}, error) {
	fmt.Printf("Cargando datos de la colección '%s' con el filtro: %v\n", s.Model.CollectionName, filter)
	data, err := db.FindDocuments(filter, s.Model.CollectionName) // Simula db.FindDocuments
	return data, err
}

// Método Insert: Insertar datos en la base de datos
func (s *BaseModelController) Insert(data map[string]interface{}) error {
	fmt.Printf("Insertando datos en la colección '%s': %v\n", s.Model.CollectionName, data)
	// Lógica real de inserción en la base de datos aquí.
	return nil // Retorna nil como ejemplo de éxito
}

// Método Update: Actualizar datos en la base de datos
func (s *BaseModelController) Update(filter map[string]interface{}, data map[string]interface{}) error {
	fmt.Printf("Actualizando datos de la colección '%s' con el filtro: %v y los datos: %v\n", s.Model.CollectionName, filter, data)
	// Lógica real de actualización aquí.
	return nil // Retorna nil como ejemplo de éxito
}

// Método Delete: Eliminar datos en la base de datos
func (s *BaseModelController) Delete(filter map[string]interface{}) error {
	fmt.Printf("Eliminando datos de la colección '%s' con el filtro: %v\n", s.Model.CollectionName, filter)
	// Lógica real de eliminación aquí.
	return nil // Retorna nil como ejemplo de éxito
}

var controllers []map[string]*BaseModelController

// Crear un nuevo controlador
func NewController(Model *Model) BaseModelController {
	newController := &BaseModelController{Model: Model}
	saveController(map[string]*BaseModelController{Model.Name: newController})
	return *newController
}

func saveController(newController map[string]*BaseModelController) {
	controllers = append(controllers, newController)
}

func SaveController(controller *BaseModelController) {
	existingController := GetController(controller.Model.Name)
	if existingController != nil {
		for i, controllerMap := range controllers {
			if _, exists := controllerMap[controller.Model.Name]; exists {
				controllers[i][controller.Model.Name] = controller
				return
			}
		}
	} else {
		saveController(map[string]*BaseModelController{controller.Model.Name: controller})
	}
}

// Obtener todos los controladores
func GetControllers() []*BaseModelController {
	var result []*BaseModelController
	for _, controllerMap := range controllers {
		for _, controller := range controllerMap {
			result = append(result, controller)
		}
	}
	return result
}

// Obtener un controlador por nombre de modelo
func GetController(name string) *BaseModelController {
	for _, controllerMap := range controllers {
		for _, controller := range controllerMap {
			if controller.Model.Name == name {
				return controller
			}
		}
	}
	return nil
}

var methods []map[string]Methods

func SaveMethods(newMethods map[string]Methods) {
	methods = append(methods, newMethods)
}
