package gson_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	gson "github.com/DiscoFighter47/gSON"
	"github.com/stretchr/testify/assert"
)

func TestServeData(t *testing.T) {
	testData := []struct {
		des  string
		body interface{}
	}{
		{"serve struct", &body{"Hello Universe!"}},
		{"serve object", gson.Object{"msg": "Hello Universe!"}},
	}
	for _, td := range testData {
		t.Run(td.des, func(t *testing.T) {
			r := httptest.NewRecorder()
			gson.ServeData(r, td.body)
			assert.Equal(t, http.StatusOK, r.Code)
			assert.JSONEq(t, `{"data":{"msg":"Hello Universe!"}}`, string(r.Body.Bytes()))
		})
	}
}
