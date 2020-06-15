package rest_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ipoval/goactivesupport/rest"
	"github.com/stretchr/testify/assert"
)

func TestRenderJsonErr(t *testing.T) {
	t.Run("given error should render the error", func(t *testing.T) {
		err := errors.New("missing header")
		respRec := httptest.NewRecorder()
		rest.RenderJsonErr(respRec, http.StatusBadRequest, err)
		assert.Contains(t, respRec.Body.String(), `{"error":"missing header"}`)
		assert.Equal(t, respRec.Header().Get("Content-Type"), rest.MimeType.Json)
	})

	t.Run("given no error should render default error", func(t *testing.T) {
		respRec := httptest.NewRecorder()
		rest.RenderJsonErr(respRec, http.StatusBadRequest, nil)
		assert.Contains(t, respRec.Body.String(), `{"error":"bad request"}`)
		assert.Equal(t, respRec.Header().Get("Content-Type"), rest.MimeType.Json)
	})
}
