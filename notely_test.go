package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addParseTimeParam(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{"ValidURL", "mysql://user:pass@host/db", "mysql://user:pass@host/db?parseTime=true", nil},
		{"InvalidURL", "not-a-url", "", nil},
		{"MissingScheme", "user:pass@host/db", "http://user:pass@host/db?parseTime=true", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := addParseTimeParam(tc.input)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_handlerReadiness(t *testing.T) {
	req := httptest.NewRequest("GET", "/v1/healthz", nil)
	rr := httptest.NewRecorder()

	handlerReadiness(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "{\"status\":\"ok\"}", rr.Body.String())
}
