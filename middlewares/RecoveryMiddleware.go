package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/hrrydgls/snug/exceptions"
	"github.com/hrrydgls/snug/models/responses"
)

func RecoveryMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		defer func() {
			if rec := recover(); rec != nil {
				var code = http.StatusInternalServerError
				var message = "Internal Server Error!"

				if apiError, ok := rec.(exceptions.APIError); ok {
					code = apiError.Code
					message = apiError.Message
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(code)

				json.NewEncoder(w).Encode(responses.Response{
					Message: message,
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}