package user

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ssshekhu53/user-detail-management/errors"
	"github.com/ssshekhu53/user-detail-management/models"
	"github.com/ssshekhu53/user-detail-management/utils"
)

func Test_Create(t *testing.T) {
	u := New().(*user)

	userReq := &models.User{
		Fname:   "John",
		City:    "New York",
		Phone:   "1234567890",
		Height:  5.9,
		Married: false,
	}

	id := u.Create(userReq)

	assert.Equal(t, int(1), id)
	assert.Equal(t, u.users[id].Fname, userReq.Fname)
	assert.Equal(t, u.users[id].City, userReq.City)
	assert.Equal(t, u.users[id].Phone, userReq.Phone)
	assert.Equal(t, u.users[id].Height, userReq.Height)
	assert.Equal(t, u.users[id].Married, userReq.Married)
}

func Test_Get(t *testing.T) {
	u := New().(*user)
	userReq1 := &models.User{Fname: "John", City: "New York", Phone: "1234567890", Height: 5.9, Married: false}
	userReq2 := &models.User{Fname: "Jane", City: "San Francisco", Phone: "0987654321", Height: 5.5, Married: true}

	u.Create(userReq1)
	u.Create(userReq2)

	tests := []struct {
		name    string
		filters *models.Filters
		want    []models.User
	}{
		{"Get all users", nil, []models.User{u.users[1], u.users[2]}},
		{"Filter by City", &models.Filters{City: utils.StrPtr("New York")}, []models.User{u.users[1]}},
		{"Filter by Married status", &models.Filters{Married: utils.BoolPtr(true)}, []models.User{u.users[2]}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := u.Get(tt.filters)

			assert.ElementsMatch(t, tt.want, users)
		})
	}
}

func Test_GetByID(t *testing.T) {
	u := New().(*user)
	userReq := &models.User{Fname: "John", City: "New York", Phone: "1234567890", Height: 5.9, Married: false}
	id := u.Create(userReq)

	usr, err := u.GetByID(id)
	assert.NoError(t, err)
	assert.Equal(t, u.users[id], *usr)

	_, err = u.GetByID(999)

	assert.Error(t, err)
	assert.IsType(t, errors.UserNotFound{}, err)
}

func Test_Update(t *testing.T) {
	u := New().(*user)
	usrCreateReq := &models.User{Fname: "John", City: "New York", Phone: "1234567890", Height: 5.9, Married: false}
	id := u.Create(usrCreateReq)

	usrUpdateReq := &models.User{ID: id, Fname: "Johnny", City: "Los Angeles", Phone: "0987654321", Height: 6.0, Married: true}

	updatedUser := &models.User{ID: id, Fname: "Johnny", City: "Los Angeles", Phone: "0987654321", Height: 6.0, Married: true}
	u.Update(usrUpdateReq)

	usr := u.users[id]

	assert.Equal(t, *updatedUser, usr)
}

func Test_Delete(t *testing.T) {
	u := New().(*user)
	userReq := &models.User{Fname: "John", City: "New York", Phone: "1234567890", Height: 5.9, Married: false}
	id := u.Create(userReq)

	u.Delete(id)

	usr, ok := u.users[id]

	assert.Empty(t, usr)
	assert.False(t, ok)
}

func Test_isMatch(t *testing.T) {
	u := New().(*user)
	usr := &models.User{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 5.9, Married: false}

	tests := []struct {
		name    string
		filters *models.Filters
		want    bool
	}{
		{"Match by Fname", &models.Filters{Fname: utils.StrPtr("John")}, true},
		{"Mismatch by Fname", &models.Filters{Fname: utils.StrPtr("Jane")}, false},
		{"Match by City", &models.Filters{City: utils.StrPtr("New York")}, true},
		{"Mismatch by City", &models.Filters{City: utils.StrPtr("San Francisco")}, false},
		{"Match by Height", &models.Filters{Height: utils.Float64Ptr(5.9)}, true},
		{"Mismatch by Height", &models.Filters{Height: utils.Float64Ptr(6.0)}, false},
		{"Match by Married status", &models.Filters{Married: utils.BoolPtr(false)}, true},
		{"Mismatch by Married status", &models.Filters{Married: utils.BoolPtr(true)}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := u.isMatch(usr, tt.filters)

			assert.Equal(t, tt.want, got)
		})
	}
}
