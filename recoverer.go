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
					serveError(w, err)
				case error:
					log.Println(string(debug.Stack()))
					serveError(w, NewAPIerror("Internal Server Error", http.StatusInternalServerError, err))
				case string:
					log.Println(string(debug.Stack()))
					serveError(w, NewAPIerror("Internal Server Error", http.StatusInternalServerError, errors.New(err)))
				default:
					panic(err)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func serveError(w http.ResponseWriter, err *APIerror) {
	res := Response{
		Status: err.Status,
		Error:  err,
	}
	res.ServeJSON(w)
}
