// Archivo generado automáticamente para el módulo unidadMedida (model)
package unidad_medida_model

import (
	base_models "main/pkg/base/models"
)

type UnidadMedidaStruct struct {
	Id     *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre string `db:"nombre"`
}

var Model = base_models.NewModel[UnidadMedidaStruct]("unidadMedida", "unidadMedida")
