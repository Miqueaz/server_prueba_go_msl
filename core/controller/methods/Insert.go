package methods

import "fmt"

// Implementa el método Insert
func (s *BaseMethods) Insert(data map[string]interface{}, collectionName string) {
	fmt.Printf("Insertando datos en la colección '%s': %v\n", collectionName, data)
	// Aquí puedes agregar la lógica real, como acceder a una base de datos.
}
