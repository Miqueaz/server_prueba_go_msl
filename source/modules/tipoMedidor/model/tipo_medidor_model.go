// Archivo generado automáticamente para el módulo tipoMedidor (model)
package tipo_medidor_model

import (
	base_models "main/pkg/base/models"
)

type TipoMedidorStruct struct {
	Id          *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre      string `db:"nombre"`
	Descripcion string `db:"descripcion"`
}

var Model = base_models.NewModel[TipoMedidorStruct]("tipoMedidor", "tipoMedidor")
