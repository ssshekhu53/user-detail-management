package models

type Filters struct {
	Fname   *string  `json:"fname"`
	City    *string  `json:"city"`
	Height  *float64 `json:"height"`
	Married *bool    `json:"married"`
}
