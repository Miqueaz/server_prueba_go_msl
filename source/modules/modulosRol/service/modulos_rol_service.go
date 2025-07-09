package modulos_rol_service

import (
	modulos_rol_model "main/source/modules/modulosRol/model"
	modulos_model "main/source/modules/modulos/model"
	modulos_service "main/source/modules/modulos/service"
	roles_service "main/source/modules/roles/service"
	base_service "main/pkg/base/service"
)

type ModulosRolService struct {
	base_service.Service[modulos_rol_model.ModulosRolStruct]
}

func (s *ModulosRolService) Read() ([]modulos_rol_model.ModulosRolSanitizer, error) {
	var err error
	var roles []modulos_rol_model.ModulosRolSanitizer
	var modulos []modulos_model.ModulosStruct
	var relations []modulos_rol_model.ModulosRolStruct
	
	if roles, err = roles_service.Service.Read(); err != nil {
		return nil, fmt.Errorf("error al leer roles: %w", err)
	}

	if modulos, err = modulos_service.Service.Read(); err != nil {
		return nil, fmt.Errorf("error al leer módulos: %w", err)
	}

	if relations, err = s.Service.Read(); err != nil {
		return nil, fmt.Errorf("error al leer relaciones módulo-rol: %w", err)
	}

	mapModulos := make(map[string]modulos_model.ModulosStruct)
	for _, modulo := range modulos {
		mapModulos[modulo.ID] = modulo
	}
	
	mapRelations := make(map[string][]modulos_model.ModulosStruct)
	for _, relation := range relations {
		if modulo, exists := mapModulos[relation.ModuloID]; exists {
			mapRelations[relation.RoleID] = append(mapRelations[relation.RoleID], modulo)
		}
	}

	dataSanitizer := make([]modulos_rol_model.ModulosRolSanitizer, 0, len(roles))
	for _, role := range roles {
		roleName := role.Nombre // para obtener dirección de una variable concreta
		dataSanitizer = append(dataSanitizer, modulos_rol_model.ModulosRolSanitizer{
			Role:    &roleName,
			Modulos: mapRelations[role.ID],
		})
	}

	return dataSanitizer, nil
}



var Service = base_service.NewService[ModulosRolService](*modulos_rol_model.Model)
