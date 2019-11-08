package gson

import (
	"encoding/json"
	"net/http"

	gero "github.com/DiscoFighter47/gEro"
)

// Object ...
type Object map[string]interface{}

// Response ...
type Response struct {
	Status int            `json:"-"`
	Data   interface{}    `json:"data,omitempty"`
	Error  *gero.APIerror `json:"error,omitempty"`
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

// ServeData ...
func ServeData(w http.ResponseWriter, data interface{}) {
	res := Response{
		Status: http.StatusOK,
		Data:   data,
	}
	res.ServeJSON(w)
}

// ServeError ...
func ServeError(w http.ResponseWriter, err *gero.APIerror) {
	res := Response{
		Status: err.Status,
		Error:  err,
	}
	res.ServeJSON(w)
}
