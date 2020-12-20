package apiserver_test

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/apiserver"
	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestHandlers_Me(t *testing.T) {
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
			expectedBody: "\"email\":\"example@dev.env\"",
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

		req, _ := http.NewRequest("GET", "/users/me", bytes.NewReader([]byte{}))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", tc.auth) // Should bind json return error otherwise

		srv.ServeHTTP(w, req)

		assert.Equal(t, tc.expectedCode, w.Code)
		body := w.Body.String()
		assert.Contains(t, body, tc.expectedBody)
	}
}

func TestServer_UserCreate(t *testing.T) {

	testCases := []struct {
		expectedCode int
		data         map[string]interface{}
		response     string
	}{
		{
			expectedCode: 201,
			data: map[string]interface{}{
				"email":                 "myemail@example.com",
				"password":              "qwerty123",
				"password_confirmation": "qwerty123",
			},
			response: "\"email\":\"myemail@example.com\"",
		},
		{
			expectedCode: 500,
			data: map[string]interface{}{
				"email":                 "myemail@example.com",
				"password":              "qwerty123",
				"password_confirmation": "qwerty123",
			},
			response: "Could not save the user",
		},
		{
			expectedCode: 422,
			data: map[string]interface{}{
				"email":                 "",
				"password":              "qwerty123",
				"password_confirmation": "qwerty123",
			},
			response: "\"Email\":\"Email must be valid Email\"",
		},
		{
			expectedCode: 422,
			data: map[string]interface{}{
				"email":                 "myemail@example.com",
				"password":              "qwerty123",
				"password_confirmation": "",
			},
			response: "\"PasswordConfirmation\":\"Passwords does not match\"",
		},
	}
	srv := apiserver.TestServer(t)

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		reqJson, err := json.Marshal(tc.data)
		if err != nil {
			panic(err)
		}

		req, _ := http.NewRequest("POST", "/users/", bytes.NewReader(reqJson))
		req.Header.Add("Content-Type", "application/json") // Should bind json return error otherwise

		srv.ServeHTTP(w, req)

		assert.Equal(t, tc.expectedCode, w.Code)
		body := w.Body.String()
		assert.Contains(t, body, tc.response)
	}
}
func TestServer_UserCreateValidationErrors(t *testing.T) {

	testCases := []struct {
		expectedCode int
		data         map[string]interface{}
		response     map[string]map[string]string
	}{
		{
			expectedCode: 422,
			data: map[string]interface{}{
				"email":                 "nonemail.com",
				"password":              "qwerty123",
				"password_confirmation": "",
			},
			response: map[string]map[string]string{
				"error": {
					"Email": "Email must be valid Email",
				},
			},
		},
		{
			expectedCode: 422,
			data: map[string]interface{}{
				"email":                 "nonemailmecom",
				"password":              "passwr",
				"password_confirmation": "",
			},
			response: map[string]map[string]string{
				"error": {
					"Password": "Password is too short, accepted length: 8",
					"Email":    "Email must be valid Email",
				},
			},
		},
	}

	srv := apiserver.TestServer(t)

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		reqJson, err := json.Marshal(tc.data)
		if err != nil {
			panic(err)
		}

		req, _ := http.NewRequest("POST", "/users/", bytes.NewReader(reqJson))
		req.Header.Add("Content-Type", "application/json") // Should bind json return error otherwise

		srv.ServeHTTP(w, req)

		assert.Equal(t, tc.expectedCode, w.Code)
		body := w.Body.String()
		for _, msg := range tc.response["error"] {
			assert.Contains(t, body, msg)
		}
	}
}

func TestServer_DeckCreate(t *testing.T) {

	testCases := []struct {
		expectedCode int
		data         string
		auth         bool
		response     string
	}{
		{
			expectedCode: 201,
			data:         "NEWDECK",
			auth:         true,
			response:     "\"name\":\"NEWDECK\"",
		},
	}
	srv := apiserver.TestServer(t)

	// Creating and Retrieving user
	userModel := &models.User{}
	{
		user := struct {
			expectedCode int
			data         map[string]interface{}
			response     string
		}{
			expectedCode: 201,
			data: map[string]interface{}{
				"email":                 "myemail@example.com",
				"password":              "qwerty123",
				"password_confirmation": "qwerty123",
			},
			response: "\"email\":\"myemail@example.com\"",
		}
		w := httptest.NewRecorder()
		reqJson, err := json.Marshal(user.data)
		if err != nil {
			panic(err)
		}
		req, _ := http.NewRequest("POST", "/users/", bytes.NewReader(reqJson))
		req.Header.Add("Content-Type", "application/json") // Should bind json return error otherwise
		srv.ServeHTTP(w, req)
		assert.Equal(t, user.expectedCode, w.Code)
		body := w.Body
		err = json.Unmarshal(body.Bytes(), userModel)
		if err != nil {
			panic(err)
		}
	}
	for _, tc := range testCases {
		wData := map[string]string{
			"name": tc.data,
		}
		w := httptest.NewRecorder()
		reqJson, err := json.Marshal(wData)
		if err != nil {
			panic(err)
		}

		req, _ := http.NewRequest("POST", "/deck/", bytes.NewReader(reqJson))
		req.Header.Add("Content-Type", "application/json") // Should bind json return error otherwise
		if tc.auth {
			req.Header.Add("Authorization", authHeader(t, "myemail@example.com", "qwerty123"))
		}

		srv.ServeHTTP(w, req)

		assert.Equal(t, tc.expectedCode, w.Code)
		body := w.Body.String()
		assert.Contains(t, body, tc.response)
	}
}

func authHeader(t *testing.T, email string, password string) string {
	t.Helper()

	sEnc := b64.StdEncoding.EncodeToString([]byte(email + ":" + password))

	return "Basic " + sEnc
}
