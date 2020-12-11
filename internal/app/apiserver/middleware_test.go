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

func TestServer_TestAuth(t *testing.T) {
	userToCreate := map[string]string{
		"email":                 "example@dev.env",
		"password":              "mypassword",
		"password_confirmation": "mypassword",
	}
	testCases := []struct {
		auth         string
		expectedCode int
		expectedBody string
	}{
		{
			expectedCode: http.StatusOK,
			auth:         "Basic ZXhhbXBsZUBkZXYuZW52Om15cGFzc3dvcmQ=",
			expectedBody: "\"email\"",
		},
		{
			expectedCode: http.StatusUnauthorized,
			auth:         "",
			expectedBody: "\"error\":\"Unauthorized\"",
		},
	}

	srv := apiserver.TestServer(t)

	{ // Create user
		w := httptest.NewRecorder()
		reqJson, err := json.Marshal(userToCreate)
		if err != nil {
			panic(err)
		}
		req, _ := http.NewRequest("POST", "/users/", bytes.NewReader(reqJson))
		req.Header.Add("Content-Type", "application/json") // Should bind json return error otherwise
		srv.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)
	}

	for _, tc := range testCases {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/users/", bytes.NewReader([]byte{}))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", tc.auth) // Should bind json return error otherwise

		srv.ServeHTTP(w, req)

		assert.Equal(t, tc.expectedCode, w.Code)
		body := w.Body.String()
		assert.Contains(t, body, tc.expectedBody)
	}
}
func TestServer_OnlyJSON(t *testing.T) {

	testCases := []struct {
		expectedCode int
		headers      map[string]string
		expectedBody string
	}{
		{
			expectedCode: http.StatusOK,
			headers:      map[string]string{"Content-Type": "application/json"},
			expectedBody: "Testing is ok",
		},
		{
			expectedCode: http.StatusUnsupportedMediaType,
			headers:      map[string]string{"Content-Type": "plain/text"},
			expectedBody: "\"Unsupported content type\"",
		},
	}

	srv := apiserver.TestServer(t)

	for _, tc := range testCases {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/test/", bytes.NewReader([]byte{}))
		for hKey, hVal := range tc.headers {
			req.Header.Add(hKey, hVal)
		}
		srv.ServeHTTP(w, req)

		assert.Equal(t, tc.expectedCode, w.Code)
		body := w.Body.String()
		assert.Contains(t, body, tc.expectedBody)
	}
}
