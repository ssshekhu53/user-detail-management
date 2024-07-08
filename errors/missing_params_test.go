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
			wantErr: "missing params",
		},
		{
			name:    "One missing param",
			params:  []string{"phone"},
			wantErr: "missing param: phone",
		},
		{
			name:    "Multiple missing params",
			params:  []string{"phone", "height"},
			wantErr: "missing params: phone, height",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MissingParams{Params: tt.params}
			assert.EqualError(t, err, tt.wantErr)
		})
	}
}
