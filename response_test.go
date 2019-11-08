package gson_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	gero "github.com/DiscoFighter47/gEro"

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

func TestServeError(t *testing.T) {
	r := httptest.NewRecorder()
	gson.ServeError(r, gero.NewAPIerror("error", http.StatusInternalServerError, fmt.Errorf("demo error"), "tag"))
	assert.Equal(t, http.StatusInternalServerError, r.Code)
	assert.JSONEq(t, `{"error":{"title":"error","detail":"demo error","tags":["tag"]}}`, string(r.Body.Bytes()))
}
