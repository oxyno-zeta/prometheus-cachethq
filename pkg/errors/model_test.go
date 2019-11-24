package errors

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewNotFoundError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *GeneralError
	}{
		{
			name: "Not found error case",
			args: args{
				err: errors.New("Not found"),
			},
			want: &GeneralError{
				ErrorType: NotFoundErrorType,
				Err:       errors.New("Not found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotFoundError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInternalServerError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *GeneralError
	}{
		{
			name: "Internal server error case",
			args: args{
				err: errors.New("Internal server"),
			},
			want: &GeneralError{
				ErrorType: InternalServerErrorType,
				Err:       errors.New("Internal server"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInternalServerError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternalServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBadInputError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *GeneralError
	}{
		{
			name: "Bad input error case",
			args: args{
				err: errors.New("Bad input"),
			},
			want: &GeneralError{
				ErrorType: BadInputErrorType,
				Err:       errors.New("Bad input"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBadInputError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBadInputError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneralError_Error(t *testing.T) {
	type fields struct {
		ErrorType string
		Err       error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Error 1 case",
			fields: fields{
				ErrorType: "Type1",
				Err:       errors.New("Error1"),
			},
			want: "Error1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ge := &GeneralError{
				ErrorType: tt.fields.ErrorType,
				Err:       tt.fields.Err,
			}
			if got := ge.Error(); got != tt.want {
				t.Errorf("GeneralError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
