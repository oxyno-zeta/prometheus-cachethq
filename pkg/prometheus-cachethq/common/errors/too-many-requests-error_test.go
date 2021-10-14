// +build unit

package errors

import (
	gerrors "errors"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestNewTooManyRequestsError(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name       string
		args       args
		err        error
		ext        map[string]interface{}
		statusCode int
	}{
		{
			name:       "constructor",
			args:       args{msg: "fake"},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS"},
			statusCode: 429,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTooManyRequestsError(tt.args.msg)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewTooManyRequestsError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewTooManyRequestsError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewTooManyRequestsError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewTooManyRequestsError().stackTrace must exists")
			}
		})
	}
}

func TestNewTooManyRequestsErrorWithError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name       string
		args       args
		err        error
		ext        map[string]interface{}
		statusCode int
	}{
		{
			name:       "constructor",
			args:       args{err: errors.New("fake")},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS"},
			statusCode: 429,
		},
		{
			name:       "constructor with golang error",
			args:       args{err: gerrors.New("fake")},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS"},
			statusCode: 429,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTooManyRequestsErrorWithError(tt.args.err)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewTooManyRequestsErrorWithError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewTooManyRequestsErrorWithError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewTooManyRequestsErrorWithError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewTooManyRequestsErrorWithError().stackTrace must exists")
			}
		})
	}
}

func TestNewTooManyRequestsErrorWithExtensions(t *testing.T) {
	type args struct {
		msg              string
		customExtensions map[string]interface{}
	}
	tests := []struct {
		name       string
		args       args
		err        error
		ext        map[string]interface{}
		statusCode int
	}{
		{
			name: "constructor with nil map",
			args: args{
				msg: "fake",
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS"},
			statusCode: 429,
		},
		{
			name: "constructor with empty map",
			args: args{
				msg:              "fake",
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS"},
			statusCode: 429,
		},
		{
			name: "constructor with existing map",
			args: args{
				msg: "fake",
				customExtensions: map[string]interface{}{
					"fake": 1,
				},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS", "fake": 1},
			statusCode: 429,
		},
		{
			name: "constructor with override map",
			args: args{
				msg: "fake",
				customExtensions: map[string]interface{}{
					"code": 1,
					"test": true,
				},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS", "test": true},
			statusCode: 429,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTooManyRequestsErrorWithExtensions(tt.args.msg, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewTooManyRequestsErrorWithExtensions().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewTooManyRequestsErrorWithExtensions().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewTooManyRequestsErrorWithExtensions().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewTooManyRequestsErrorWithExtensions().stackTrace must exists")
			}
		})
	}
}

func TestNewTooManyRequestsErrorWithExtensionsAndError(t *testing.T) {
	type args struct {
		err              error
		customExtensions map[string]interface{}
	}
	tests := []struct {
		name       string
		args       args
		err        error
		ext        map[string]interface{}
		statusCode int
	}{
		{
			name: "constructor with nil map",
			args: args{
				err: errors.New("fake"),
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS"},
			statusCode: 429,
		},
		{
			name: "constructor with empty map",
			args: args{
				err:              errors.New("fake"),
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS"},
			statusCode: 429,
		},
		{
			name: "constructor with existing map",
			args: args{
				err: errors.New("fake"),
				customExtensions: map[string]interface{}{
					"fake": 1,
				},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS", "fake": 1},
			statusCode: 429,
		},
		{
			name: "constructor with override map",
			args: args{
				err: errors.New("fake"),
				customExtensions: map[string]interface{}{
					"code": 1,
					"test": true,
				},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS", "test": true},
			statusCode: 429,
		},
		{
			name: "constructor with golang error",
			args: args{
				err: gerrors.New("fake"),
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "TOO_MANY_REQUESTS"},
			statusCode: 429,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTooManyRequestsErrorWithExtensionsAndError(tt.args.err, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewTooManyRequestsErrorWithExtensionsAndError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewTooManyRequestsErrorWithExtensionsAndError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewTooManyRequestsErrorWithExtensionsAndError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewTooManyRequestsErrorWithExtensionsAndError().stackTrace must exists")
			}
		})
	}
}
