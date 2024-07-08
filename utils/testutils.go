package utils

// Helper functions to create pointers for primitive types

func IntPtr(i int) *int {
	return &i
}

func StrPtr(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

func Float64Ptr(f float64) *float64 {
	if f == 0 {
		return nil
	}

	return &f
}

func BoolPtr(b bool) *bool {
	return &b
}
