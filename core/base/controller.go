package base

// Controller agrupa un modelo con sus métodos
type Controller struct {
	Model
	Methods
}

// Mapa global para almacenar controladores
var controllers = make(map[string]*Controller)

// Crear un nuevo controlador con métodos por defecto
func NewController(model Model) {
	methods := &BaseModelController{Model: model}
	controller := &Controller{Model: model, Methods: methods}
	SaveController(controller)
}

// Crear un nuevo controlador con métodos personalizados
func NewControllerWithMethods(model Model, methods Methods) {
	controller := &Controller{Model: model, Methods: methods}
	SaveController(controller)
}

// Guardar un controlador en el mapa global
func SaveController(controller *Controller) {
	controllers[controller.Model.Name] = controller
}

// Obtener un controlador por su nombre
func GetController(name string) *Controller {
	return controllers[name]
}
