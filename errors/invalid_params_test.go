package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidParamsError(t *testing.T) {
	tests := []struct {
		name    string
		params  []string
		wantErr string
	}{
		{
			name:    "No invalid params",
			params:  []string{},
			wantErr: "Invalid Params",
		},
		{
			name:    "One invalid param",
			params:  []string{"phone"},
			wantErr: "Invalid Param: phone",
		},
		{
			name:    "Multiple invalid params",
			params:  []string{"phone", "height"},
			wantErr: "Invalid Params: phone, height",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InvalidParams{Params: tt.params}
			assert.EqualError(t, err, tt.wantErr)
		})
	}
}
