package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingParamsError(t *testing.T) {
	tests := []struct {
		name    string
		params  []string
		wantErr string
	}{
		{
			name:    "No missing params",
			params:  []string{},
			wantErr: "Missing Params",
		},
		{
			name:    "One missing param",
			params:  []string{"phone"},
			wantErr: "Missing Param: phone",
		},
		{
			name:    "Multiple missing params",
			params:  []string{"phone", "height"},
			wantErr: "Missing Params: phone, height",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MissingParams{Params: tt.params}
			assert.EqualError(t, err, tt.wantErr)
		})
	}
}
