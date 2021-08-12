// +build unit

package errors

import (
	gerrors "errors"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestNewNotFoundError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "NOT_FOUND"},
			statusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNotFoundError(tt.args.msg)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewNotFoundError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewNotFoundError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewNotFoundError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewNotFoundError().stackTrace must exists")
			}
		})
	}
}

func TestNewNotFoundErrorWithError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "NOT_FOUND"},
			statusCode: 404,
		},
		{
			name:       "constructor with golang error",
			args:       args{err: gerrors.New("fake")},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "NOT_FOUND"},
			statusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNotFoundErrorWithError(tt.args.err)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewNotFoundErrorWithError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewNotFoundErrorWithError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewNotFoundErrorWithError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewNotFoundErrorWithError().stackTrace must exists")
			}
		})
	}
}

func TestNewNotFoundErrorWithExtensions(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "NOT_FOUND"},
			statusCode: 404,
		},
		{
			name: "constructor with empty map",
			args: args{
				msg:              "fake",
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "NOT_FOUND"},
			statusCode: 404,
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
			ext:        map[string]interface{}{"code": "NOT_FOUND", "fake": 1},
			statusCode: 404,
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
			ext:        map[string]interface{}{"code": "NOT_FOUND", "test": true},
			statusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNotFoundErrorWithExtensions(tt.args.msg, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewNotFoundErrorWithExtensions().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewNotFoundErrorWithExtensions().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewNotFoundErrorWithExtensions().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewNotFoundErrorWithExtensions().stackTrace must exists")
			}
		})
	}
}

func TestNewNotFoundErrorWithExtensionsAndError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "NOT_FOUND"},
			statusCode: 404,
		},
		{
			name: "constructor with empty map",
			args: args{
				err:              errors.New("fake"),
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "NOT_FOUND"},
			statusCode: 404,
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
			ext:        map[string]interface{}{"code": "NOT_FOUND", "fake": 1},
			statusCode: 404,
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
			ext:        map[string]interface{}{"code": "NOT_FOUND", "test": true},
			statusCode: 404,
		},
		{
			name: "constructor with golang error",
			args: args{
				err: gerrors.New("fake"),
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "NOT_FOUND"},
			statusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNotFoundErrorWithExtensionsAndError(tt.args.err, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewNotFoundErrorWithExtensionsAndError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewNotFoundErrorWithExtensionsAndError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewNotFoundErrorWithExtensionsAndError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewNotFoundErrorWithExtensionsAndError().stackTrace must exists")
			}
		})
	}
}
