// Archivo generado automáticamente para el módulo medidores (routes)
package medidores_routes

import (
	"main/modules/core/router"
	medidores_core "main/modules/medidores/core"
	"net/http"
)

func Init() {
	var r = router.Router()
	r.GET("/medidores", http.HandlerFunc(medidores_core.Controller.Read))
}
