// Archivo generado automáticamente para el módulo metas (model)
package metas_model

import (
	base_models "main/pkg/base/models"
)

type MetasStruct struct {
	Id          *int   `db:"id" sanitizer:"id" visible:"false"`
	Titulo      string `db:"titulo"`
	Descripcion string `db:"descripcion"`
	Industria   *int   `db:"industria"`
}

var Model = base_models.NewModel[MetasStruct]("metas", "metas")
