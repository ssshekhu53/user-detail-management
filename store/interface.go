package store

import "github.com/ssshekhu53/user-detail-management/models"

//go:generate mockgen -source=interface.go -destination=mock_interface.go -package=store

type User interface {
	Create(user *models.User) int
	Get(filters *models.Filters) []models.User
	GetByID(id int) (*models.User, error)
	Update(user *models.User)
	Delete(id int)
}
