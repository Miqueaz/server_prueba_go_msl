package helpers

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// Función para convertir BSON a `map[string]interface{}` y manejar `primitive.ObjectID`
func ConvertBsonToMap(doc bson.M) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range doc {
		result[key] = fmt.Sprintf("%v", value)
	}
	return result
}

func ConvertMapToStrings(input map[string]interface{}) (map[string]string, error) {
	result := make(map[string]string)
	for key, value := range input {
		strValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("el valor para la clave %s no es un string", key)
		}
		result[key] = strValue
	}
	return result, nil
}
