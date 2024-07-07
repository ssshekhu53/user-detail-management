package user

import (
	"sort"
	"strings"

	"github.com/ssshekhu53/user-detail-management/errors"
	"github.com/ssshekhu53/user-detail-management/models"
	"github.com/ssshekhu53/user-detail-management/store"
)

type user struct {
	users          map[int]models.User
	lastInsertedID int
}

func New() store.User {
	u := &user{}
	u.users = make(map[int]models.User)
	u.lastInsertedID = 0

	return u
}

func (u *user) Create(userReq *models.User) int {
	u.lastInsertedID += 1

	userReq.ID = u.lastInsertedID

	u.users[u.lastInsertedID] = *userReq

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

func (u *user) GetByID(id int) (*models.User, error) {
	if usr, ok := u.users[id]; ok {
		return &usr, nil
	}

	return nil, errors.UserNotFound{ID: id}
}

func (u *user) GetByIDs(ids []int) []models.User {
	users := make([]models.User, 0)

	idsMap := make(map[int]bool)

	for _, id := range ids {
		idsMap[id] = true
	}

	for _, usr := range u.users {
		if _, ok := idsMap[usr.ID]; ok {
			users = append(users, usr)
		}
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	return users
}

func (u *user) Update(usr *models.User) {
	u.users[usr.ID] = *usr
}

func (u *user) Delete(id int) {
	delete(u.users, id)
}

func (u *user) isMatch(usr *models.User, filters *models.Filters) bool {
	if filters.Fname != nil && !strings.EqualFold(usr.Fname, *filters.Fname) {
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
