// Archivo generado automáticamente para el módulo estadoActividad (model)
package estado_actividad_model

import (
	base_models "main/pkg/base/models"
)

type EstadoActividadStruct struct {
	Id     *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre string `db:"nombre"`
}

var Model = base_models.NewModel[EstadoActividadStruct]("estadoActividad", "estadoActividad")
