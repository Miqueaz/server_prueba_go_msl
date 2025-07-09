// Archivo generado automáticamente para el módulo archivos (model)
package archivos_model

import (
	base_models "main/pkg/base/models"
	"time"
)

type ArchivosStruct struct {
	Id	*int	`db:"id" sanitizer:"id" visible:"false"`
	Nombre	string	`db:"nombre"`
	Url	string	`db:"url"`
	Creado	*time.Time	`db:"creado"`
	Tipo	*int	`db:"tipo"`
}

var Model = base_models.NewModel[ArchivosStruct]("archivos", "archivos")
