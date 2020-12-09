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
