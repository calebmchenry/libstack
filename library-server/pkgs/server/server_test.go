package server_test

import (
	"bytes"
	"encoding/json"
	"libstack/pkgs/server"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s server.Server

func TestMain(m *testing.M) {
	s = server.New()
	code := m.Run()
	os.Exit(code)
}

func TestLogin(t *testing.T) {
	t.Run("returns token", func(t *testing.T) {
		var jsonStr = []byte("fake")
		req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		res := Do(req)
		assert.Equal(t, http.StatusOK, res.Code)

		var payload server.Payload[interface{}]
		err := json.NewDecoder(res.Body).Decode(&payload)
		if err != nil {
			t.Errorf("Decoding of json response failed for login: %s", err.Error())
		}
		assert.Empty(t, payload.Error)
		assert.True(t, payload.Ok)
		assert.NotEmpty(t, payload.Data)
	})
	t.Run("only allows application/json", func(t *testing.T) {
		var jsonStr = []byte("fake")
		req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "foo")
		res := Do(req)
		assert.Equal(t, http.StatusUnsupportedMediaType, res.Code)
	})
	t.Run("not ok when called with invalid creds", func(t *testing.T) {
		// TODO(mchenryc): handle when creds are actually checked
	})
}

func TestLogout(t *testing.T) {
	t.Run("forbids logging out of unauthenticated users", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/api/v1/logout", nil)
		res := Do(req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})
	// TODO(mchenryc): test success logout response
}

func Do(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}
