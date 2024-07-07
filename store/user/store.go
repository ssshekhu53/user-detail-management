package user

import (
	"sort"
	"strings"

	"github.com/ssshekhu53/user-detail-management/errors"
	"github.com/ssshekhu53/user-detail-management/models"
	"github.com/ssshekhu53/user-detail-management/store"
)

type user struct {
	users          map[int64]models.User
	lastInsertedID int64
}

func New() store.User {
	u := &user{}
	u.users = make(map[int64]models.User)
	u.lastInsertedID = 0

	return u
}

func (u *user) Create(userReq *models.UserRequest) int64 {
	u.lastInsertedID += 1

	u.users[u.lastInsertedID] = models.User{
		ID:      u.lastInsertedID,
		FName:   *userReq.FName,
		City:    *userReq.City,
		Phone:   *userReq.Phone,
		Height:  *userReq.Height,
		Married: *userReq.Married,
	}

	return u.lastInsertedID
}

func (u *user) Get(filters *models.Filters) []models.User {
	users := make([]models.User, 0)

	for _, usr := range u.users {
		if filters == nil || u.isMatch(&usr, filters) {
			users = append(users, usr)
		}
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	return users
}

func (u *user) GetByID(id int64) (*models.User, error) {
	if usr, ok := u.users[id]; ok {
		return &usr, nil
	}

	return nil, errors.UserNotFound{ID: id}
}

func (u *user) Update(user *models.User) {
	u.users[user.ID] = *user
}

func (u *user) Delete(id int64) {
	delete(u.users, id)
}

func (u *user) isMatch(usr *models.User, filters *models.Filters) bool {
	if filters.Fname != nil && !strings.EqualFold(usr.FName, *filters.Fname) {
		return false
	}

	if filters.City != nil && !strings.EqualFold(usr.City, *filters.City) {
		return false
	}

	if filters.Height != nil && usr.Height != *filters.Height {
		return false
	}

	if filters.Married != nil && usr.Married != *filters.Married {
		return false
	}

	return true
}
