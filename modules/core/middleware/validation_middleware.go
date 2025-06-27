package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ctxKey string

const DtoKey ctxKey = "dto"

var validate = validator.New()

func ValidatorMiddleware[T any](next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto T
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			sendJSON(w, "Invalid JSON: "+err.Error(), 400)
			return
		}

		if err := validate.Struct(dto); err != nil {
			sendJSON(w, "Validation error: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := runCustomValidator(&dto); err != nil {
			sendJSON(w, "Custom validation error: "+err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), DtoKey, &dto)
		next(w, r.WithContext(ctx))
	}
}

func runCustomValidator(dto interface{}) error {
	v := reflect.ValueOf(dto)
	t := reflect.TypeOf(dto)

	//Execute every methods to dto type
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)

		if strings.HasPrefix(method.Name, "Valid") {
			result := v.MethodByName(method.Name).Call(nil)
			if len(result) != 1 {
				return errors.New("method" + method.Name + " should return a error")
			}
			if err, ok := result[0].Interface().(error); ok && err != nil {
				return err
			}
		}
	}
	return nil
}
