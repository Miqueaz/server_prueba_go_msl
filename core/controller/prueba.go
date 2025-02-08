package main

import "fmt"

// BaseModelController simula tu estructura base con un método predeterminado
type BaseModelController struct {
	caja Caja
}

type Caja struct {
	Name string
}

// Método predeterminado para GetName
func (b BaseModelController) GetName() string {
	return "Base: " + b.caja.Name
}

// Define una interfaz común para todos los modelos
type Model interface {
	GetName() string
}

// commentModel es una estructura que incluye BaseModelController
type commentModel struct {
	BaseModelController
}

// Sobrescribe el método GetName para commentModel
func (c commentModel) GetName() string {
	return "Comment: " + c.caja.Name
}

// postModel es otra estructura que incluye BaseModelController
type postModel struct {
	BaseModelController
}

// postModel usa el método GetName predeterminado de BaseModelController

func main() {
	// Crea instancias de las estructuras
	comment := commentModel{
		BaseModelController: BaseModelController{caja: Caja{Name: "CommentModel"}},
	}
	post := postModel{
		BaseModelController: BaseModelController{caja: Caja{Name: "PostModel"}},
	}

	// Crea un slice de la interfaz Model para almacenar ambos tipos
	var models []Model
	print(comment.caja.Name)
	models = append(models, comment, post)

	// Itera sobre los modelos y utiliza los métodos comunes
	for _, model := range models {
		fmt.Println(model.GetName()) // Salida: "Comment: CommentModel", "Base: PostModel"
	}
}
