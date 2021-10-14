// +build unit

package errors

import (
	gerrors "errors"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestNewUnauthorizedError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "UNAUTHORIZED"},
			statusCode: 401,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUnauthorizedError(tt.args.msg)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewUnauthorizedError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewUnauthorizedError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewUnauthorizedError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewUnauthorizedError().stackTrace must exists")
			}
		})
	}
}

func TestNewUnauthorizedErrorWithError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "UNAUTHORIZED"},
			statusCode: 401,
		},
		{
			name:       "constructor with golang error",
			args:       args{err: gerrors.New("fake")},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "UNAUTHORIZED"},
			statusCode: 401,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUnauthorizedErrorWithError(tt.args.err)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewUnauthorizedErrorWithError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewUnauthorizedErrorWithError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewUnauthorizedErrorWithError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewUnauthorizedErrorWithError().stackTrace must exists")
			}
		})
	}
}

func TestNewUnauthorizedErrorWithExtensions(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "UNAUTHORIZED"},
			statusCode: 401,
		},
		{
			name: "constructor with empty map",
			args: args{
				msg:              "fake",
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "UNAUTHORIZED"},
			statusCode: 401,
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
			ext:        map[string]interface{}{"code": "UNAUTHORIZED", "fake": 1},
			statusCode: 401,
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
			ext:        map[string]interface{}{"code": "UNAUTHORIZED", "test": true},
			statusCode: 401,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUnauthorizedErrorWithExtensions(tt.args.msg, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewUnauthorizedErrorWithExtensions().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewUnauthorizedErrorWithExtensions().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewUnauthorizedErrorWithExtensions().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewUnauthorizedErrorWithExtensions().stackTrace must exists")
			}
		})
	}
}

func TestNewUnauthorizedErrorWithExtensionsAndError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "UNAUTHORIZED"},
			statusCode: 401,
		},
		{
			name: "constructor with empty map",
			args: args{
				err:              errors.New("fake"),
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "UNAUTHORIZED"},
			statusCode: 401,
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
			ext:        map[string]interface{}{"code": "UNAUTHORIZED", "fake": 1},
			statusCode: 401,
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
			ext:        map[string]interface{}{"code": "UNAUTHORIZED", "test": true},
			statusCode: 401,
		},
		{
			name: "constructor with golang error",
			args: args{
				err: gerrors.New("fake"),
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "UNAUTHORIZED"},
			statusCode: 401,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUnauthorizedErrorWithExtensionsAndError(tt.args.err, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewUnauthorizedErrorWithExtensionsAndError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewUnauthorizedErrorWithExtensionsAndError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewUnauthorizedErrorWithExtensionsAndError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewUnauthorizedErrorWithExtensionsAndError().stackTrace must exists")
			}
		})
	}
}
