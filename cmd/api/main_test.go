package main_test

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	m "github.com/brettearle/galf/cmd/api"
	tu "github.com/brettearle/galf/internal/testutil"
)

func Test(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)
	go m.Run(ctx, m.Config{
		Host: "0.0.0.0",
		Port: "8080",
	}, os.Stderr)

	// Time for server ready
	tu.WaitForReady(ctx, 2*time.Second, "http://0.0.0.0:8080/api/health")

	//Header Values
	jsonContent := "application/json"
	textContent := "text/plain"

	cases := []struct {
		name           string
		endpoint       string
		method         string
		body           string
		contentType    *string
		expectedStatus int
	}{
		{
			name:           "tt: /health returns 200",
			method:         http.MethodGet,
			body:           "",
			endpoint:       "http://0.0.0.0:8080/api/health",
			contentType:    nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "tt: /register post with correct data returns 200",
			method:         http.MethodPost,
			body:           `{"name":"test","state":"on"}`,
			endpoint:       "http://0.0.0.0:8080/api/register",
			contentType:    &jsonContent,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "tt: /register GET should fail",
			method:         http.MethodGet,
			body:           `{"name":"not under test","state":"on"}`,
			endpoint:       "http://0.0.0.0:8080/api/register",
			contentType:    &jsonContent,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "tt: /register PUT should fail",
			method:         http.MethodPut,
			body:           `{"name":"not under test","state":"on"}`,
			endpoint:       "http://0.0.0.0:8080/api/register",
			contentType:    &jsonContent,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "tt: /register PATCH should fail",
			method:         http.MethodPatch,
			body:           `{"name":"not under test","state":"on"}`,
			endpoint:       "http://0.0.0.0:8080/api/register",
			contentType:    &jsonContent,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "tt: /register DELETE should fail",
			method:         http.MethodDelete,
			body:           `{"name":"not under test","state":"on"}`,
			endpoint:       "http://0.0.0.0:8080/api/register",
			contentType:    &jsonContent,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "tt: /register post with incorrect data 'empty name' returns 422",
			method:         http.MethodPost,
			body:           `{"name":"","state":"on"}`,
			endpoint:       "http://0.0.0.0:8080/api/register",
			contentType:    &jsonContent,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:           "tt: /register post with incorrect data 'wrong state string' returns 422",
			method:         http.MethodPost,
			body:           `{"name":"name is good","state":"wrong"}`,
			endpoint:       "http://0.0.0.0:8080/api/register",
			contentType:    &jsonContent,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:           "tt: /register post with incorrect content type header returns",
			method:         http.MethodPost,
			body:           `can't decode this`,
			endpoint:       "http://0.0.0.0:8080/api/register",
			contentType:    &textContent,
			expectedStatus: http.StatusUnsupportedMediaType,
		},
		{
			name:           "tt: /flag/test1 with correct name returns 200",
			method:         http.MethodGet,
			body:           "",
			endpoint:       "http://0.0.0.0:8080/api/flag/test1",
			contentType:    nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "tt: /flag with correct name returns 200",
			method:         http.MethodGet,
			body:           "",
			endpoint:       "http://0.0.0.0:8080/api/flag/test1",
			contentType:    nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "tt: /flag returns 200",
			method:         http.MethodGet,
			body:           "",
			endpoint:       "http://0.0.0.0:8080/api/flag",
			contentType:    nil,
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			body := bytes.NewBufferString(tt.body)

			req, err := http.NewRequestWithContext(ctx, tt.method, tt.endpoint, body)
			if err != nil {
				t.Errorf("Failed to create request")
			}

			if tt.contentType != nil {
				req.Header.Add("Content-Type", *tt.contentType)
			}

			client := http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Error sending request")
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf(tu.Red("got %d want %d"), resp.StatusCode, http.StatusOK)
			}
		})
	}

}
