package gson_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	gson "github.com/DiscoFighter47/gSON"
	"github.com/stretchr/testify/assert"
)

func TestServeData(t *testing.T) {
	t.Run("serve struct", func(t *testing.T) {
		b := &body{"Hello Universe!"}
		r := httptest.NewRecorder()
		gson.ServeData(r, b)
		assert.Equal(t, http.StatusOK, r.Code)
		assert.JSONEq(t, `{"data":{"msg":"Hello Universe!"}}`, string(r.Body.Bytes()))
	})

	t.Run("serve object", func(t *testing.T) {
		r := httptest.NewRecorder()
		gson.ServeData(r, gson.Object{"msg": "Hello Universe!"})
		assert.Equal(t, http.StatusOK, r.Code)
		assert.JSONEq(t, `{"data":{"msg":"Hello Universe!"}}`, string(r.Body.Bytes()))
	})
}
