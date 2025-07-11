// Archivo generado automáticamente para el módulo tipoActividad (model)
package tipo_actividad_model

import (
	base_models "main/pkg/base/models"
)

type TipoActividadStruct struct {
	Id     *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre string `db:"nombre"`
}

var Model = base_models.NewModel[TipoActividadStruct]("tipoActividad", "tipoActividad")
