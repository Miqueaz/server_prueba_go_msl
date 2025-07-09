package base_controller

import (
	"encoding/json"
	"errors"
	"io"
	"main/pkg/client"
	"reflect"
	"strconv"

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

	return func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}

		argsCount := fnType.NumIn()
		args := make([]reflect.Value, argsCount)
		stringParamIndex := 0

		for i := 0; i < argsCount; i++ {
			paramType := fnType.In(i)

			if paramType == reflect.TypeOf((*gin.Context)(nil)) {
				args[i] = reflect.ValueOf(c)
				continue
			}

			switch paramType.Kind() {
			case reflect.Struct:
				paramPtr := reflect.New(paramType)
				if len(bodyBytes) > 0 {
					if err := json.Unmarshal(bodyBytes, paramPtr.Interface()); err != nil {
						c.JSON(400, gin.H{"error": "invalid request body (struct): " + err.Error()})
						return
					}
				}
				args[i] = paramPtr.Elem()

			case reflect.Map:
				paramPtr := reflect.New(paramType).Interface()
				if len(bodyBytes) > 0 {
					if err := json.Unmarshal(bodyBytes, paramPtr); err != nil {
						c.JSON(400, gin.H{"error": "invalid request body (map): " + err.Error()})
						return
					}
				}
				args[i] = reflect.ValueOf(paramPtr).Elem()

			case reflect.String:
				if stringParamIndex < len(c.Params) {
					args[i] = reflect.ValueOf(c.Params[stringParamIndex].Value)
				} else {
					args[i] = reflect.Zero(paramType)
				}
				stringParamIndex++

			case reflect.Int:
				if stringParamIndex < len(c.Params) {
					val := c.Params[stringParamIndex].Value
					if intValue, err := strconv.Atoi(val); err == nil {
						args[i] = reflect.ValueOf(intValue)
					}
				}
				stringParamIndex++

			case reflect.Float64:
				if stringParamIndex < len(c.Params) {
					val := c.Params[stringParamIndex].Value
					if floatValue, err := strconv.ParseFloat(val, 64); err == nil {
						args[i] = reflect.ValueOf(floatValue)
					}
				}
				stringParamIndex++

			default:
				args[i] = reflect.Zero(paramType)
			}
		}

		results := fnVal.Call(args)

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

		if fnErr != nil {
			c.JSON(500, gin.H{"error": fnErr.Error()})
			return
		}
		c.JSON(200, resp)
	}
}
