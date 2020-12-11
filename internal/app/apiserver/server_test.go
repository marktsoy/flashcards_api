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

func TestServer_UserCreate(t *testing.T) {

	testCases := []struct {
		expectedCode int
		data         map[string]interface{}
		response     map[string]interface{}
	}{
		{
			expectedCode: 201,
			data: map[string]interface{}{
				"email":                 "myemail@example.com",
				"password":              "qwerty123",
				"password_confirmation": "qwerty123",
			},
			response: map[string]interface{}{
				"email": "myemail@example.com",
			},
		},
		{
			expectedCode: 500,
			data: map[string]interface{}{
				"email":                 "myemail@example.com",
				"password":              "qwerty123",
				"password_confirmation": "qwerty123",
			},
			response: map[string]interface{}{
				"error": "Could not save the user",
			},
		},
		{
			expectedCode: 422,
			data: map[string]interface{}{
				"email":                 "",
				"password":              "qwerty123",
				"password_confirmation": "qwerty123",
			},
			response: map[string]interface{}{
				"error": "Invalid request",
			},
		},
		{
			expectedCode: 422,
			data: map[string]interface{}{
				"email":                 "myemail@example.com",
				"password":              "qwerty123",
				"password_confirmation": "",
			},
			response: map[string]interface{}{
				"error": "Invalid request",
			},
		},
	}

	type res struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}

	type errRes struct {
		Error string `json:"error"`
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
		body := w.Body.Bytes()
		switch tc.expectedCode {
		default:
			{
				res := &errRes{}
				jsonParseError := json.Unmarshal(body, res)
				assert.NoError(t, jsonParseError)
				assert.Equal(t, tc.response["error"], res.Error)
			}
		case 201:
			res := &res{}
			jsonParseError := json.Unmarshal(body, res)
			assert.NoError(t, jsonParseError)
			assert.Equal(t, tc.response["email"], res.Email)
		}
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
