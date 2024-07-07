package user

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ssshekhu53/user-detail-management/errors"
	"github.com/ssshekhu53/user-detail-management/models"
)

func TestCreate(t *testing.T) {
	u := New().(*user)

	userReq := &models.UserRequest{
		FName:   strPtr("John"),
		City:    strPtr("New York"),
		Phone:   strPtr("1234567890"),
		Height:  float64Ptr(5.9),
		Married: boolPtr(false),
	}

	id := u.Create(userReq)

	assert.Equal(t, int(1), id)
	assert.Equal(t, u.users[id].FName, *userReq.FName)
	assert.Equal(t, u.users[id].City, *userReq.City)
	assert.Equal(t, u.users[id].Phone, *userReq.Phone)
	assert.Equal(t, u.users[id].Height, *userReq.Height)
	assert.Equal(t, u.users[id].Married, *userReq.Married)
}

func TestGet(t *testing.T) {
	u := New().(*user)
	userReq1 := &models.UserRequest{FName: strPtr("John"), City: strPtr("New York"), Phone: strPtr("1234567890"), Height: float64Ptr(5.9), Married: boolPtr(false)}
	userReq2 := &models.UserRequest{FName: strPtr("Jane"), City: strPtr("San Francisco"), Phone: strPtr("0987654321"), Height: float64Ptr(5.5), Married: boolPtr(true)}

	u.Create(userReq1)
	u.Create(userReq2)

	tests := []struct {
		name    string
		filters *models.Filters
		want    []models.User
	}{
		{"Get all users", nil, []models.User{u.users[1], u.users[2]}},
		{"Filter by City", &models.Filters{City: strPtr("New York")}, []models.User{u.users[1]}},
		{"Filter by Married status", &models.Filters{Married: boolPtr(true)}, []models.User{u.users[2]}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := u.Get(tt.filters)

			assert.ElementsMatch(t, tt.want, users)
		})
	}
}

func TestGetByID(t *testing.T) {
	u := New().(*user)
	userReq := &models.UserRequest{FName: strPtr("John"), City: strPtr("New York"), Phone: strPtr("1234567890"), Height: float64Ptr(5.9), Married: boolPtr(false)}
	id := u.Create(userReq)

	usr, err := u.GetByID(id)
	assert.NoError(t, err)
	assert.Equal(t, u.users[id], *usr)

	_, err = u.GetByID(999)

	assert.Error(t, err)
	assert.IsType(t, errors.UserNotFound{}, err)
}

func TestUpdate(t *testing.T) {
	u := New().(*user)
	userReq := &models.UserRequest{FName: strPtr("John"), City: strPtr("New York"), Phone: strPtr("1234567890"), Height: float64Ptr(5.9), Married: boolPtr(false)}
	id := u.Create(userReq)

	updatedUser := &models.User{ID: id, FName: "Johnny", City: "Los Angeles", Phone: "0987654321", Height: 6.0, Married: true}
	u.Update(updatedUser)

	usr := u.users[id]

	assert.Equal(t, *updatedUser, usr)
}

func TestDelete(t *testing.T) {
	u := New().(*user)
	userReq := &models.UserRequest{FName: strPtr("John"), City: strPtr("New York"), Phone: strPtr("1234567890"), Height: float64Ptr(5.9), Married: boolPtr(false)}
	id := u.Create(userReq)

	u.Delete(id)

	usr, ok := u.users[id]

	assert.Empty(t, usr)
	assert.False(t, ok)
}

func TestIsMatch(t *testing.T) {
	u := New().(*user)
	usr := &models.User{ID: 1, FName: "John", City: "New York", Phone: "1234567890", Height: 5.9, Married: false}

	tests := []struct {
		name    string
		filters *models.Filters
		want    bool
	}{
		{"Match by FName", &models.Filters{Fname: strPtr("John")}, true},
		{"Mismatch by FName", &models.Filters{Fname: strPtr("Jane")}, false},
		{"Match by City", &models.Filters{City: strPtr("New York")}, true},
		{"Mismatch by City", &models.Filters{City: strPtr("San Francisco")}, false},
		{"Match by Height", &models.Filters{Height: float64Ptr(5.9)}, true},
		{"Mismatch by Height", &models.Filters{Height: float64Ptr(6.0)}, false},
		{"Match by Married status", &models.Filters{Married: boolPtr(false)}, true},
		{"Mismatch by Married status", &models.Filters{Married: boolPtr(true)}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := u.isMatch(usr, tt.filters)

			assert.Equal(t, tt.want, got)
		})
	}
}

// Helper functions to create pointers for primitive types
func strPtr(s string) *string {
	return &s
}

func float64Ptr(f float64) *float64 {
	return &f
}

func boolPtr(b bool) *bool {
	return &b
}
