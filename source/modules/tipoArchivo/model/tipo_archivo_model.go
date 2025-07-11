// Archivo generado automáticamente para el módulo tipoArchivo (model)
package tipo_archivo_model

import (
	base_models "main/pkg/base/models"
)

type TipoArchivoStruct struct {
	Id     *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre string `db:"nombre"`
}

var Model = base_models.NewModel[TipoArchivoStruct]("tipoArchivo", "tipoArchivo")
