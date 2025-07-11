// Archivo generado automáticamente para el módulo lineasAgua (model)
package lineas_agua_model

import (
	base_models "main/pkg/base/models"
)

type LineasAguaStruct struct {
	Id      *int   `db:"id" sanitizer:"id" visible:"false"`
	Codigo  string `db:"codigo"`
	Area    *int   `db:"area"`
	Medidor *int   `db:"medidor"`
}

var Model = base_models.NewModel[LineasAguaStruct]("lineasAgua", "lineasAgua")
