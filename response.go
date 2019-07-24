package gson

import (
	"encoding/json"
	"net/http"
)

// Object ...
type Object map[string]interface{}

// Response ...
type Response struct {
	Code int         `json:"-"`
	Data interface{} `json:"data,omitempty"`
}

// ServeJSON ...
func (res *Response) ServeJSON(w http.ResponseWriter) {
	if res.Code == 0 {
		res.Code = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
