package gson

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"
)

// Recoverer ...
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error Occurred:", err)
				switch err := err.(type) {
				case *APIerror:
					res := Response{
						Status: err.Status,
						Error:  err,
					}
					res.ServeJSON(w)
				case error:
					res := Response{
						Status: http.StatusInternalServerError,
						Error:  NewAPIerror("Internal Server Error", http.StatusInternalServerError, err),
					}
					res.ServeJSON(w)
					log.Println(string(debug.Stack()))
				case string:
					res := Response{
						Status: http.StatusInternalServerError,
						Error:  NewAPIerror("Internal Server Error", http.StatusInternalServerError, errors.New(err)),
					}
					res.ServeJSON(w)
					log.Println(string(debug.Stack()))
				default:
					panic(err)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
