// Archivo generado automáticamente para el módulo fuenteLectura (model)
package fuente_lectura_model

import (
	base_models "main/pkg/base/models"
)

type FuenteLecturaStruct struct {
	Id          *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre      string `db:"nombre"`
	Descripcion string `db:"descripcion"`
}

var Model = base_models.NewModel[FuenteLecturaStruct]("fuenteLectura", "fuenteLectura")
