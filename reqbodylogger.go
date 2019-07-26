package gson

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ReqBodyLogger ...
func ReqBodyLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			buf, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(NewAPIerror("Unable To Read Request Body", http.StatusUnprocessableEntity, err))
			}
			var objmap map[string]*json.RawMessage
			if err = json.Unmarshal(buf, &objmap); err == nil {
				b, err := json.Marshal(objmap)
				if err == nil {
					log.Println("Handling Request:", string(b))
				}
			}
			r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		}
		next.ServeHTTP(w, r)
	})
}
