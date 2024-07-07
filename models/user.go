package models

import "github.com/ssshekhu53/user-detail-management/errors"

type User struct {
	ID      int     `json:"id"`
	FName   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   string  `json:"phone"`
	Height  float64 `json:"height"`
	Married bool    `json:"married"`
}

type UserRequest struct {
	FName   *string  `json:"fname"`
	City    *string  `json:"city"`
	Phone   *string  `json:"phone"`
	Height  *float64 `json:"height"`
	Married *bool    `json:"married"`
}

func (u UserRequest) ValidateMissingParam() error {
	var missing []string

	if u.FName == nil {
		missing = append(missing, "fname")
	}

	if u.City == nil {
		missing = append(missing, "city")
	}

	if u.Phone == nil {
		missing = append(missing, "phone")
	}

	if u.Height == nil {
		missing = append(missing, "height")
	}

	if u.Married == nil {
		missing = append(missing, "married")
	}

	if len(missing) > 0 {
		return errors.MissingParams{Params: missing}
	}

	return nil
}

func (u UserRequest) ValidateInvalidParam() error {
	var invalid []string

	if len(*u.Phone) < 10 {
		invalid = append(invalid, "phone")
	}

	if *u.Height <= 0.0 {
		invalid = append(invalid, "height")
	}

	if len(invalid) > 0 {
		return errors.InvalidParams{Params: invalid}
	}

	return nil
}

type UserUpdateRequest struct {
	ID      *int     `json:"id"`
	FName   *string  `json:"fname"`
	City    *string  `json:"city"`
	Phone   *string  `json:"phone"`
	Height  *float64 `json:"height"`
	Married *bool    `json:"married"`
}

func (u UserUpdateRequest) ValidateMissingParam() error {
	var missing []string

	if u.ID == nil {
		missing = append(missing, "id")
	}

	if u.FName == nil {
		missing = append(missing, "fname")
	}

	if u.City == nil {
		missing = append(missing, "city")
	}

	if u.Phone == nil {
		missing = append(missing, "phone")
	}

	if u.Height == nil {
		missing = append(missing, "height")
	}

	if u.Married == nil {
		missing = append(missing, "married")
	}

	if len(missing) > 0 {
		return errors.MissingParams{Params: missing}
	}

	return nil
}

func (u UserUpdateRequest) ValidateInvalidParam() error {
	var invalid []string

	if *u.ID <= 0 {
		invalid = append(invalid, "id")
	}

	if len(*u.Phone) < 10 {
		invalid = append(invalid, "phone")
	}

	if *u.Height <= 0.0 {
		invalid = append(invalid, "height")
	}

	if len(invalid) > 0 {
		return errors.InvalidParams{Params: invalid}
	}

	return nil
}
