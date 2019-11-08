package gson_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	gson "github.com/DiscoFighter47/gSON"
)

type body struct {
	Msg string `json:"msg"`
}

func TestParseBody(t *testing.T) {
	t.Run("parse body", func(t *testing.T) {
		b := &body{}
		r := httptest.NewRequest(http.MethodGet, "https://google.com", strings.NewReader(`{"msg":"Hello Universe!"}`))
		assert.NoError(t, gson.ParseBody(r, b))
		assert.Equal(t, "Hello Universe!", b.Msg)
	})

	t.Run("innvalid body", func(t *testing.T) {
		b := &body{}
		r := httptest.NewRequest(http.MethodGet, "https://google.com", strings.NewReader(``))
		assert.Error(t, gson.ParseBody(r, b))
	})
}
