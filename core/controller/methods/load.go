package methods

import (
	"fmt"
	"server/core/connection/db"
)

// Implementa el método Load
func (s *BaseMethods) Load(filter map[string]interface{}, collectionName string) (interface{}, error) {
	fmt.Printf("Cargando datos de la colección '%s' con el filtro: %v\n", collectionName, filter)
	data, err := db.FindDocuments(filter, collectionName)
	return data, err
}

// func Load() *structure.BaseController {
// 	return &structure.BaseController{
// 		Load: func(filter map[string]interface{}, collectionName string) {
// 			// Implementation here
// 		},
// 	}
// }
