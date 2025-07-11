// Archivo generado automáticamente para el módulo tipoAlerta (model)
package tipo_alerta_model

import (
	base_models "main/pkg/base/models"
)

type TipoAlertaStruct struct {
	Id     *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre string `db:"nombre"`
}

var Model = base_models.NewModel[TipoAlertaStruct]("tipoAlerta", "tipoAlerta")
