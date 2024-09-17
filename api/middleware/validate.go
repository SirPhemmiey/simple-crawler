package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRequest(i interface{}) func(http.HandlerFunc) http.HandlerFunc {
	t := reflect.TypeOf(i)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Create a new instance of the struct
			request := reflect.New(t).Interface()

			// Decode the request into the struct
			err := json.NewDecoder(r.Body).Decode(request)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Validate the struct
			if err := validate.Struct(request); err != nil {
				http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
				return
			}

			// Store the validated request in the context
			ctx := context.WithValue(r.Context(), "validatedRequest", request)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
