// Archivo generado automáticamente para el módulo lecturas (model)
package lecturas_model

import (
	base_models "main/pkg/base/models"
	"time"
)

type LecturasStruct struct {
	Id	*int	`db:"id" sanitizer:"id" visible:"false"`
	Valor	*float64	`db:"valor"`
	FechaRegistro	*time.Time	`db:"fecharegistro"`
	FuenteLectura	*int	`db:"fuentelectura"`
}

var Model = base_models.NewModel[LecturasStruct]("lecturas", "lecturas")
