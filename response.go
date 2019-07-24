package gson

import (
	"encoding/json"
	"net/http"
)

// Object ...
type Object map[string]interface{}

// Response ...
type Response struct {
	Status int         `json:"-"`
	Data   interface{} `json:"data,omitempty"`
	Error  *APIerror   `json:"error,omitempty"`
}

// ServeJSON ...
func (res *Response) ServeJSON(w http.ResponseWriter) {
	if res.Status == 0 {
		res.Status = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
