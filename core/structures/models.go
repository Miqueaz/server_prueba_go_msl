package Structures

type Model struct {
	Name           string      `json:"name"`
	CollectionName string      `json:"collectionName"`
	Structure      interface{} `json:"structure"`
}

var models = []map[string]*Model{}

func NewModel(name string, collectionName string, structure interface{}) (*Model, error) {
	model := &Model{
		Name:           name,
		CollectionName: collectionName,
		Structure:      structure,
	}
	// Guardar el modelo (se asume que esta función maneja cualquier error internamente)
	saveModel(model)
	return model, nil
}

func GetModels() []*Model {
	var result []*Model
	for _, modelMap := range models {
		for _, model := range modelMap {
			result = append(result, model)
		}
	}
	return result
}

func GetModel(name string) *Model {
	for _, model := range models {
		if m, ok := model[name]; ok {
			return m
		}
	}
	return nil
}

func saveModel(model *Model) {
	models = append(models, map[string]*Model{model.Name: model})
}
