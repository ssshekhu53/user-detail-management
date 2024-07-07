package user

import (
	"github.com/ssshekhu53/user-detail-management/errors"
	"github.com/ssshekhu53/user-detail-management/models"
	"github.com/ssshekhu53/user-detail-management/service"
	"github.com/ssshekhu53/user-detail-management/store"
)

type user struct {
	userStore store.User
}

func New(userStore store.User) service.User {
	return &user{userStore: userStore}
}

func (u *user) Create(usr *models.UserRequest) (*models.User, error) {
	existingUsers := u.Search(&models.Filters{Fname: usr.Fname, City: usr.City, Phone: usr.Phone, Height: usr.Height, Married: usr.Married})
	if len(existingUsers) != 0 {
		return nil, errors.UserAlreadyExists{}
	}

	newUser := &models.User{
		Fname:   *usr.Fname,
		City:    *usr.City,
		Phone:   *usr.Phone,
		Height:  *usr.Height,
		Married: *usr.Married,
	}

	id := u.userStore.Create(newUser)

	newUser, _ = u.userStore.GetByID(id)

	return newUser, nil
}

func (u *user) Get() []models.User {
	users := u.userStore.Get(nil)

	return users
}

func (u *user) GetByID(id int) (*models.User, error) {
	usr, err := u.userStore.GetByID(id)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *user) Update(usr *models.UserUpdateRequest) (*models.User, error) {
	existingUser, err := u.userStore.GetByID(*usr.ID)
	if err != nil {
		return nil, err
	}

	existingUser.Fname = *usr.Fname
	existingUser.City = *usr.City
	existingUser.Phone = *usr.Phone
	existingUser.Height = *usr.Height
	existingUser.Married = *usr.Married

	u.userStore.Update(existingUser)

	updatedUser, _ := u.userStore.GetByID(*usr.ID)

	return updatedUser, nil
}

func (u *user) Delete(id int) error {
	_, err := u.userStore.GetByID(id)
	if err != nil {
		return err
	}

	u.userStore.Delete(id)

	return nil
}

func (u *user) Search(filters *models.Filters) []models.User {
	users := u.userStore.Get(filters)

	return users
}
