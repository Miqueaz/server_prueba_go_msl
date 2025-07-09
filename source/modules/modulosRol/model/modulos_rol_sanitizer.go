package modulos_rol_sanitizer 

import (
	modulos "main/source/modules/modulos/model"
)

type ModulosRoleSanitizer struct {
	Role *string
	Modulos []modulos.ModulosStruct
}
