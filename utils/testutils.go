package utils

// Helper functions to create pointers for primitive types

func IntPtr(i int) *int {
	return &i
}

func StrPtr(s string) *string {
	return &s
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func BoolPtr(b bool) *bool {
	return &b
}
