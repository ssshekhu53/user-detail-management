package models

type Filters struct {
	Fname  *string  `json:"fname"`
	City   *string  `json:"city"`
	Phone  *string  `json:"phone"`
	Height *float64 `json:"height"`
}
