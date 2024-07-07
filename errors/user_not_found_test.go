package errors

import "testing"

func TestUserNotFoundError(t *testing.T) {
	tests := []struct {
		name     string
		err      UserNotFound
		expected string
	}{
		{
			name:     "User ID is zero",
			err:      UserNotFound{ID: 0},
			expected: "user not found",
		},
		{
			name:     "User ID is non-zero",
			err:      UserNotFound{ID: 123},
			expected: "user with ID 123 not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, got)
			}
		})
	}
}
