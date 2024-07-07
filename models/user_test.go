package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ssshekhu53/user-detail-management/errors"
)

func TestValidateMissingParam(t *testing.T) {
	tests := []struct {
		name    string
		request UserRequest
		wantErr error
	}{
		{
			name: "All fields present",
			request: UserRequest{
				FName:   strPtr("John"),
				City:    strPtr("New York"),
				Phone:   strPtr("1234567890"),
				Height:  float64Ptr(5.9),
				Married: boolPtr(true),
			},
			wantErr: nil,
		},
		{
			name: "Missing FName",
			request: UserRequest{
				City:    strPtr("New York"),
				Phone:   strPtr("1234567890"),
				Height:  float64Ptr(5.9),
				Married: boolPtr(true),
			},
			wantErr: errors.MissingParams{Params: []string{"fname"}},
		},
		{
			name: "Missing multiple fields",
			request: UserRequest{
				FName: strPtr("John"),
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

func TestValidateInvalidParam(t *testing.T) {
	tests := []struct {
		name    string
		request UserRequest
		wantErr error
	}{
		{
			name: "All fields valid",
			request: UserRequest{
				FName:   strPtr("John"),
				City:    strPtr("New York"),
				Phone:   strPtr("1234567890"),
				Height:  float64Ptr(5.9),
				Married: boolPtr(true),
			},
			wantErr: nil,
		},
		{
			name: "Invalid Phone",
			request: UserRequest{
				FName:   strPtr("John"),
				City:    strPtr("New York"),
				Phone:   strPtr("123"),
				Height:  float64Ptr(5.9),
				Married: boolPtr(true),
			},
			wantErr: errors.InvalidParams{Params: []string{"phone"}},
		},
		{
			name: "Invalid Height",
			request: UserRequest{
				FName:   strPtr("John"),
				City:    strPtr("New York"),
				Phone:   strPtr("1234567890"),
				Height:  float64Ptr(-1.0),
				Married: boolPtr(true),
			},
			wantErr: errors.InvalidParams{Params: []string{"height"}},
		},
		{
			name: "Invalid Phone and Height",
			request: UserRequest{
				FName:   strPtr("John"),
				City:    strPtr("New York"),
				Phone:   strPtr("123"),
				Height:  float64Ptr(-1.0),
				Married: boolPtr(true),
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

// Helper functions to create pointers from literals
func strPtr(s string) *string {
	return &s
}

func float64Ptr(f float64) *float64 {
	return &f
}

func boolPtr(b bool) *bool {
	return &b
}
