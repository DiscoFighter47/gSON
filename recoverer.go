package gson

import (
	"net/http"
)

// Recoverer ...
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch err := err.(type) {
				case error:
					res := Response{
						Code: http.StatusInternalServerError,
						Data: Object{
							"error": err.Error(),
						},
					}
					res.ServeJSON(w)
				case string:
					res := Response{
						Code: http.StatusInternalServerError,
						Data: Object{
							"error": err,
						},
					}
					res.ServeJSON(w)
				default:
					panic(err)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
