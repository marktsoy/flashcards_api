package apiserver_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/apiserver"
	"github.com/stretchr/testify/assert"
)

//TestServer_NotExactURL expects request to /users to be redirected to /users/
func TestServer_NotExactURL(t *testing.T) {
	testCases := []struct {
		expectedCode int
		data         map[string]interface{}
		response     map[string]interface{}
	}{
		{
			expectedCode: 307,
		}}

	srv := apiserver.TestServer(t)

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		reqJson, err := json.Marshal(tc.data)
		if err != nil {
			panic(err)
		}

		req, _ := http.NewRequest("POST", "/users", bytes.NewReader(reqJson))
		req.Header.Add("Content-Type", "application/json") // Should bind json return error otherwise

		srv.ServeHTTP(w, req)
		assert.Equal(t, tc.expectedCode, w.Code)
		assert.Equal(t, "/users/", w.Header().Get("Location"))
	}

}
