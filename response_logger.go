package gson

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

// ResponseLogger ...
func ResponseLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := httptest.NewRecorder()
		next.ServeHTTP(res, r)
		log.Println("Serving Response:", strings.TrimSpace(res.Body.String()))

		for k, v := range res.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(res.Code)
		res.Body.WriteTo(w)
	})
}
