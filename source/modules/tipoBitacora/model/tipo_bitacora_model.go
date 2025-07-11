// Archivo generado automáticamente para el módulo tipoBitacora (model)
package tipo_bitacora_model

import (
	base_models "main/pkg/base/models"
)

type TipoBitacoraStruct struct {
	Id     *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre string `db:"nombre"`
}

var Model = base_models.NewModel[TipoBitacoraStruct]("tipoBitacora", "tipoBitacora")
