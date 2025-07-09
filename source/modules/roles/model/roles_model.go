// Archivo generado automáticamente para el módulo roles (model)
package roles_model

import (
	base_models "main/pkg/base/models"
)

type RolesStruct struct {
	Id     *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre string `db:"nombre"`
}

var Model = base_models.NewModel[RolesStruct]("roles", "roles")
