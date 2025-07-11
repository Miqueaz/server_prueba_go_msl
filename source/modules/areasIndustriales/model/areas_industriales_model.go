// Archivo generado automáticamente para el módulo areasIndustriales (model)
package areas_industriales_model

import (
	base_models "main/pkg/base/models"
)

type AreasIndustrialesStruct struct {
	Id          *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre      string `db:"nombre"`
	Descripcion string `db:"descripcion"`
	Industria   *int   `db:"industria"`
}

var Model = base_models.NewModel[AreasIndustrialesStruct]("areasIndustriales", "areasIndustriales")
