// Archivo generado automáticamente para el módulo permisosRol (model)
package modulos_rol_model

import (
	base_models "main/pkg/base/models"
)

type ModulosRolStruct struct {
	Id      *int `db:"id" sanitizer:"id" visible:"false"`
	Rol     *int `db:"rol" sanitizer:"rol" visible:"false"`
	Permiso *int `db:"permiso"`
}

var Model = base_models.NewModel[ModulosRolStruct]("permisosRol", "permisosRol")
