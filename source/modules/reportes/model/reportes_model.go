// Archivo generado automáticamente para el módulo reportes (model)
package reportes_model

import (
	base_models "main/pkg/base/models"
	"time"
)

type ReportesStruct struct {
	Id	*int	`db:"id" sanitizer:"id" visible:"false"`
	Titulo	string	`db:"titulo"`
	Descripcion	string	`db:"descripcion"`
	Creado	*time.Time	`db:"creado"`
	Industria	*int	`db:"industria"`
	Usuario	*int	`db:"usuario"`
	Archivo	*int	`db:"archivo"`
}

var Model = base_models.NewModel[ReportesStruct]("reportes", "reportes")
