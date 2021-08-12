// +build unit

package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	gerrors "errors"

	errors2 "github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/common/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetRequestURL(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:989/fake/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	want := "http://localhost:989/fake/path"
	got := GetRequestURL(req)
	if got != want {
		t.Errorf("GetRequestURI() = %v, want %v", got, want)
	}
}

func Test_RequestHost(t *testing.T) {
	hXForwardedHost1 := http.Header{
		"X-Forwarded-Host": []string{"fake.host"},
	}
	hXForwardedHost2 := http.Header{
		"X-Forwarded-Host": []string{"fake.host:9090"},
	}
	hXForwarded := http.Header{
		"Forwarded": []string{"for=192.0.2.60;proto=http;by=203.0.113.43;host=fake.host:9090"},
	}

	tests := []struct {
		name     string
		headers  http.Header
		inputURL string
		want     string
	}{
		{
			name:     "request host",
			headers:  nil,
			inputURL: "http://request.host/",
			want:     "request.host",
		},
		{
			name:     "forwarded host",
			headers:  hXForwarded,
			inputURL: "http://fake.host:9090/",
			want:     "fake.host:9090",
		},
		{
			name:     "x-forwarded host 1",
			headers:  hXForwardedHost1,
			inputURL: "http://fake.host/",
			want:     "fake.host",
		},
		{
			name:     "x-forwarded host 2",
			headers:  hXForwardedHost2,
			inputURL: "http://fake.host:9090/",
			want:     "fake.host:9090",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", tt.inputURL, nil)
			if err != nil {
				t.Fatal(err)
			}
			if tt.headers != nil {
				req.Header = tt.headers
			}

			if got := RequestHost(req); got != tt.want {
				t.Errorf("RequestHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnswerWithError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name               string
		args               args
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "not common error",
			args: args{
				err: gerrors.New("fake"),
			},
			expectedBody:       "{\"error\":\"fake\"}",
			expectedStatusCode: 500,
		},
		{
			name: "not common error 2",
			args: args{
				err: errors2.New("fake"),
			},
			expectedBody:       "{\"error\":\"fake\"}",
			expectedStatusCode: 500,
		},
		{
			name: "common conflict error",
			args: args{
				err: errors.NewConflictError("fake"),
			},
			expectedBody:       "{\"error\":\"fake\",\"extensions\":{\"code\":\"CONFLICT\"}}",
			expectedStatusCode: 409,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			AnswerWithError(c, tt.args.err)

			// Tests
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
