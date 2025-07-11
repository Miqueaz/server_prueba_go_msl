// Archivo generado automáticamente para el módulo alertas (model)
package alertas_model

import (
	base_models "main/pkg/base/models"
)

type AlertasStruct struct {
	Id          *int   `db:"id" sanitizer:"id" visible:"false"`
	Descripcion string `db:"descripcion"`
	Tipo        *int   `db:"tipo"`
}

var Model = base_models.NewModel[AlertasStruct]("alertas", "alertas")
