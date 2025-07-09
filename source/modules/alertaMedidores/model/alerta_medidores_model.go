// Archivo generado automáticamente para el módulo alertaMedidores (model)
package alerta_medidores_model

import (
	base_models "main/pkg/base/models"
	"time"
)

type AlertaMedidoresStruct struct {
	Id	*int	`db:"id" sanitizer:"id" visible:"false"`
	Medidor	*int	`db:"medidor"`
	Alerta	*int	`db:"alerta"`
	Registro	*time.Time	`db:"registro"`
}

var Model = base_models.NewModel[AlertaMedidoresStruct]("alertaMedidores", "alertaMedidores")
