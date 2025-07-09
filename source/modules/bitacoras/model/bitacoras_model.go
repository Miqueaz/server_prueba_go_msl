// Archivo generado automáticamente para el módulo bitacoras (model)
package bitacoras_model

import (
	base_models "main/pkg/base/models"
	"time"
)

type BitacorasStruct struct {
	Id	*int	`db:"id" sanitizer:"id" visible:"false"`
	Titulo	string	`db:"titulo"`
	Descripcion	string	`db:"descripcion"`
	Creado	*time.Time	`db:"creado"`
	Actualizado	*time.Time	`db:"actualizado"`
	Tipo	*int	`db:"tipo"`
	Usuario	*int	`db:"usuario"`
}

var Model = base_models.NewModel[BitacorasStruct]("bitacoras", "bitacoras")
