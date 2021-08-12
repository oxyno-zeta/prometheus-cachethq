// +build unit

package errors

import (
	gerrors "errors"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestNewInvalidInputError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "INVALID_INPUT"},
			statusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInvalidInputError(tt.args.msg)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewInvalidInputError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewInvalidInputError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewInvalidInputError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewInvalidInputError().stackTrace must exists")
			}
		})
	}
}

func TestNewInvalidInputErrorWithError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "INVALID_INPUT"},
			statusCode: 400,
		},
		{
			name:       "constructor with golang error",
			args:       args{err: gerrors.New("fake")},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "INVALID_INPUT"},
			statusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInvalidInputErrorWithError(tt.args.err)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewInvalidInputErrorWithError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewInvalidInputErrorWithError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewInvalidInputErrorWithError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewInvalidInputErrorWithError().stackTrace must exists")
			}
		})
	}
}

func TestNewInvalidInputErrorWithExtensions(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "INVALID_INPUT"},
			statusCode: 400,
		},
		{
			name: "constructor with empty map",
			args: args{
				msg:              "fake",
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "INVALID_INPUT"},
			statusCode: 400,
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
			ext:        map[string]interface{}{"code": "INVALID_INPUT", "fake": 1},
			statusCode: 400,
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
			ext:        map[string]interface{}{"code": "INVALID_INPUT", "test": true},
			statusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInvalidInputErrorWithExtensions(tt.args.msg, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewInvalidInputErrorWithExtensions().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewInvalidInputErrorWithExtensions().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewInvalidInputErrorWithExtensions().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewInvalidInputErrorWithExtensions().stackTrace must exists")
			}
		})
	}
}

func TestNewInvalidInputErrorWithExtensionsAndError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "INVALID_INPUT"},
			statusCode: 400,
		},
		{
			name: "constructor with empty map",
			args: args{
				err:              errors.New("fake"),
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "INVALID_INPUT"},
			statusCode: 400,
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
			ext:        map[string]interface{}{"code": "INVALID_INPUT", "fake": 1},
			statusCode: 400,
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
			ext:        map[string]interface{}{"code": "INVALID_INPUT", "test": true},
			statusCode: 400,
		},
		{
			name: "constructor with golang error",
			args: args{
				err: gerrors.New("fake"),
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "INVALID_INPUT"},
			statusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInvalidInputErrorWithExtensionsAndError(tt.args.err, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewInvalidInputErrorWithExtensionsAndError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewInvalidInputErrorWithExtensionsAndError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewInvalidInputErrorWithExtensionsAndError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewInvalidInputErrorWithExtensionsAndError().stackTrace must exists")
			}
		})
	}
}
