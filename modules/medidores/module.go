// Archivo de entrada para el módulo medidores
package medidores

import (
	"fmt"
	medidores_routes "main/modules/medidores/routes"
)

func Init() {
	fmt.Println("Medidores Module Initialized")
	medidores_routes.Init()
}
