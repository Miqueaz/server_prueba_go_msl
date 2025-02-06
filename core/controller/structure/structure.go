package structure

type BaseStructure struct {
	Load   func(filter map[string]interface{}, collectionName string)
	Insert func(data map[string]interface{}, collectionName string)
	Delete func(filter map[string]interface{}, collectionName string)
	Update func(filter map[string]interface{}, data map[string]interface{}, collectionName string)
}
