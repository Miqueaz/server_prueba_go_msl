// Archivo generado automáticamente para el módulo industria (model)
package industria_model

import (
	base_models "main/pkg/base/models"
)

type IndustriaStruct struct {
	Id            *int   `db:"id" sanitizer:"id" visible:"false"`
	NombreOficial string `db:"nombreoficial"`
	Alias         string `db:"alias"`
	Status        *bool  `db:"status"`
	Admin         *int   `db:"admin"`
}

var Model = base_models.NewModel[IndustriaStruct]("industria", "industria")
