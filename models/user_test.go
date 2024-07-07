package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ssshekhu53/user-detail-management/errors"
	"github.com/ssshekhu53/user-detail-management/utils"
)

func Test_UserRequestValidateMissingParam(t *testing.T) {
	tests := []struct {
		name    string
		request UserRequest
		wantErr error
	}{
		{
			name: "All fields present",
			request: UserRequest{
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("1234567890"),
				Height:  utils.Float64Ptr(5.9),
				Married: utils.BoolPtr(true),
			},
			wantErr: nil,
		},
		{
			name: "Missing Fname",
			request: UserRequest{
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("1234567890"),
				Height:  utils.Float64Ptr(5.9),
				Married: utils.BoolPtr(true),
			},
			wantErr: errors.MissingParams{Params: []string{"fname"}},
		},
		{
			name: "Missing multiple fields",
			request: UserRequest{
				Fname: utils.StrPtr("John"),
			},
			wantErr: errors.MissingParams{Params: []string{"city", "phone", "height", "married"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.ValidateMissingParam()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_UserRequestValidateInvalidParam(t *testing.T) {
	tests := []struct {
		name    string
		request UserRequest
		wantErr error
	}{
		{
			name: "All fields valid",
			request: UserRequest{
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("1234567890"),
				Height:  utils.Float64Ptr(5.9),
				Married: utils.BoolPtr(true),
			},
			wantErr: nil,
		},
		{
			name: "Invalid Phone",
			request: UserRequest{
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("123"),
				Height:  utils.Float64Ptr(5.9),
				Married: utils.BoolPtr(true),
			},
			wantErr: errors.InvalidParams{Params: []string{"phone"}},
		},
		{
			name: "Invalid Height",
			request: UserRequest{
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("1234567890"),
				Height:  utils.Float64Ptr(-1.0),
				Married: utils.BoolPtr(true),
			},
			wantErr: errors.InvalidParams{Params: []string{"height"}},
		},
		{
			name: "Invalid Phone and Height",
			request: UserRequest{
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("123"),
				Height:  utils.Float64Ptr(-1.0),
				Married: utils.BoolPtr(true),
			},
			wantErr: errors.InvalidParams{Params: []string{"phone", "height"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.ValidateInvalidParam()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_UserUpdateRequestValidateMissingParam(t *testing.T) {
	tests := []struct {
		name    string
		request UserUpdateRequest
		wantErr error
	}{
		{
			name: "All fields present",
			request: UserUpdateRequest{
				ID:      utils.IntPtr(1),
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("1234567890"),
				Height:  utils.Float64Ptr(5.9),
				Married: utils.BoolPtr(true),
			},
			wantErr: nil,
		},
		{
			name: "Missing Fname",
			request: UserUpdateRequest{
				ID:      utils.IntPtr(1),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("1234567890"),
				Height:  utils.Float64Ptr(5.9),
				Married: utils.BoolPtr(true),
			},
			wantErr: errors.MissingParams{Params: []string{"fname"}},
		},
		{
			name: "Missing multiple fields",
			request: UserUpdateRequest{
				ID:    utils.IntPtr(1),
				Fname: utils.StrPtr("John"),
			},
			wantErr: errors.MissingParams{Params: []string{"city", "phone", "height", "married"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.ValidateMissingParam()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_UserUpdateRequestValidateInvalidParam(t *testing.T) {
	tests := []struct {
		name    string
		request UserUpdateRequest
		wantErr error
	}{
		{
			name: "All fields valid",
			request: UserUpdateRequest{
				ID:      utils.IntPtr(1),
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("1234567890"),
				Height:  utils.Float64Ptr(5.9),
				Married: utils.BoolPtr(true),
			},
			wantErr: nil,
		},
		{
			name: "Invalid ID",
			request: UserUpdateRequest{
				ID:      utils.IntPtr(0),
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("1234567890"),
				Height:  utils.Float64Ptr(5.9),
				Married: utils.BoolPtr(true),
			},
			wantErr: errors.InvalidParams{Params: []string{"id"}},
		},
		{
			name: "Invalid Phone",
			request: UserUpdateRequest{
				ID:      utils.IntPtr(1),
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("123"),
				Height:  utils.Float64Ptr(5.9),
				Married: utils.BoolPtr(true),
			},
			wantErr: errors.InvalidParams{Params: []string{"phone"}},
		},
		{
			name: "Invalid Height",
			request: UserUpdateRequest{
				ID:      utils.IntPtr(1),
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("1234567890"),
				Height:  utils.Float64Ptr(-1.0),
				Married: utils.BoolPtr(true),
			},
			wantErr: errors.InvalidParams{Params: []string{"height"}},
		},
		{
			name: "Invalid Phone and Height",
			request: UserUpdateRequest{
				ID:      utils.IntPtr(1),
				Fname:   utils.StrPtr("John"),
				City:    utils.StrPtr("New York"),
				Phone:   utils.StrPtr("123"),
				Height:  utils.Float64Ptr(-1.0),
				Married: utils.BoolPtr(true),
			},
			wantErr: errors.InvalidParams{Params: []string{"phone", "height"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.ValidateInvalidParam()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
