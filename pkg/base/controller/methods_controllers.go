package base_controller

import (
	"encoding/json"
	"errors"
	"io"
	"main/pkg/client"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

// Métodos CRUD implementados por BaseModelController

// Método Read con soporte de hooks
func (s *Controller[T]) Read(c *gin.Context) {
	ctx, err := client.Init(c.Writer, c.Request)
	if err != nil {
		ctx.InternalServerError(err)
		return
	}

	data, err := s.Service.Methods.Read(nil, nil)

	if data == nil {
		err := errors.New("No data found")
		ctx.NotFound(err)
		return
	}

	// Convert []map[string]any to []any
	result := make([]any, len(data))
	for i, v := range data {
		result[i] = v
	}

	ctx.Success("Users fetched successfully", result)
}

// Insert: Insertar datos en la base de datos
func (s *Controller[T]) Insert(c *gin.Context) {
	// Implementación del método Insert
}

// Update: Actualizar datos en la base de datos
func (s *Controller[T]) Update(c *gin.Context) {
	// Implementación del método Update
}

// Delete: Eliminar datos de la base de datos
func (s *Controller[T]) Delete(c *gin.Context) {
	// Implementación del método Delete
}

// MakeController convierte cualquier función `fn` en un manejador HTTP dinámico para gin.
func MakeController(fn interface{}) gin.HandlerFunc {
	fnVal := reflect.ValueOf(fn)
	fnType := fnVal.Type()
	if fnType.Kind() != reflect.Func {
		panic("MakeController: expected a function")
	}

	// Retorna un handler de gin que ejecuta la función dinámica
	return func(c *gin.Context) {
		// Leemos el cuerpo de la solicitud
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}

		// Preparamos los argumentos para la función
		argsCount := fnType.NumIn()
		args := make([]reflect.Value, argsCount)

		// Procesamos los parámetros de la función
		for i := 0; i < argsCount; i++ {
			paramType := fnType.In(i)

			// Si el parámetro es *gin.Context, lo pasamos directamente
			if paramType == reflect.TypeOf((*gin.Context)(nil)) {
				args[i] = reflect.ValueOf(c)
				continue
			}

			// Si el parámetro es el cuerpo, lo deserializamos
			paramPtr := reflect.New(paramType)

			// Si hay datos en el cuerpo de la solicitud
			if len(bodyBytes) > 0 {
				if err := json.Unmarshal(bodyBytes, paramPtr.Interface()); err != nil {
					c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
					return
				}
			}

			args[i] = paramPtr.Elem()
		}

		// Si hay parámetros en la URL, los procesamos
		if len(c.Params) > 0 {
			for i := 0; i < argsCount; i++ {
				paramType := fnType.In(i)
				if paramType.Kind() == reflect.String {
					paramName := strings.ToLower(fnType.In(i).Name())
					paramValue := c.Param(paramName)
					if paramValue != "" {
						args[i] = reflect.ValueOf(paramValue)
					}
				}
			}
		}

		// Llamamos a la función dinámica con los argumentos
		results := fnVal.Call(args)

		// Procesamos los resultados de la función
		var resp interface{}
		var fnErr error
		switch len(results) {
		case 1:
			resp = results[0].Interface()
		case 2:
			resp = results[0].Interface()
			if e, ok := results[1].Interface().(error); ok && e != nil {
				fnErr = e
			}
		default:
			c.JSON(500, gin.H{"error": "MakeController: function must return (T) or (T, error)"})
			return
		}

		// Si hubo un error, lo manejamos
		if fnErr != nil {
			c.JSON(500, gin.H{"error": fnErr.Error()})
			return
		}

		// Respondemos con JSON
		c.JSON(200, resp)
	}
}
