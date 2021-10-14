// +build unit

package errors

import (
	gerrors "errors"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestNewForbiddenError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "FORBIDDEN"},
			statusCode: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewForbiddenError(tt.args.msg)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewForbiddenError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewForbiddenError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewForbiddenError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewForbiddenError().stackTrace must exists")
			}
		})
	}
}

func TestNewForbiddenErrorWithError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "FORBIDDEN"},
			statusCode: 403,
		},
		{
			name:       "constructor with golang error",
			args:       args{err: gerrors.New("fake")},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "FORBIDDEN"},
			statusCode: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewForbiddenErrorWithError(tt.args.err)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewForbiddenErrorWithError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewForbiddenErrorWithError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewForbiddenErrorWithError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewForbiddenErrorWithError().stackTrace must exists")
			}
		})
	}
}

func TestNewForbiddenErrorWithExtensions(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "FORBIDDEN"},
			statusCode: 403,
		},
		{
			name: "constructor with empty map",
			args: args{
				msg:              "fake",
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "FORBIDDEN"},
			statusCode: 403,
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
			ext:        map[string]interface{}{"code": "FORBIDDEN", "fake": 1},
			statusCode: 403,
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
			ext:        map[string]interface{}{"code": "FORBIDDEN", "test": true},
			statusCode: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewForbiddenErrorWithExtensions(tt.args.msg, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewForbiddenErrorWithExtensions().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewForbiddenErrorWithExtensions().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewForbiddenErrorWithExtensions().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewForbiddenErrorWithExtensions().stackTrace must exists")
			}
		})
	}
}

func TestNewForbiddenErrorWithExtensionsAndError(t *testing.T) {
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
			ext:        map[string]interface{}{"code": "FORBIDDEN"},
			statusCode: 403,
		},
		{
			name: "constructor with empty map",
			args: args{
				err:              errors.New("fake"),
				customExtensions: map[string]interface{}{},
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "FORBIDDEN"},
			statusCode: 403,
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
			ext:        map[string]interface{}{"code": "FORBIDDEN", "fake": 1},
			statusCode: 403,
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
			ext:        map[string]interface{}{"code": "FORBIDDEN", "test": true},
			statusCode: 403,
		},
		{
			name: "constructor with golang error",
			args: args{
				err: gerrors.New("fake"),
			},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "FORBIDDEN"},
			statusCode: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewForbiddenErrorWithExtensionsAndError(tt.args.err, tt.args.customExtensions)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewForbiddenErrorWithExtensionsAndError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewForbiddenErrorWithExtensionsAndError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewForbiddenErrorWithExtensionsAndError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewForbiddenErrorWithExtensionsAndError().stackTrace must exists")
			}
		})
	}
}
