package rest_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ipoval/goactivesupport/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		require.Contains(t, respRec.Body.String(), `{"error":"bad request"}`)
		require.Equal(t, respRec.Header().Get("Content-Type"), rest.MimeType.Json)
	})
}
