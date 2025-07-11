// Archivo generado automáticamente para el módulo medidores (model)
package medidores_model

import (
	base_models "main/pkg/base/models"
)

type MedidoresStruct struct {
	Id           *int   `db:"id" sanitizer:"id" visible:"false"`
	Codigo       string `db:"codigo"`
	Enlace       string `db:"enlace"`
	Tipo         *int   `db:"tipo"`
	UnidadMedida *int   `db:"unidadmedida"`
	Lecturas     *int   `db:"lecturas"`
}

var Model = base_models.NewModel[MedidoresStruct]("medidores", "medidores")
