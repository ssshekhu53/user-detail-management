package service

import "github.com/ssshekhu53/user-detail-management/models"

//go:generate mockgen -source=interface.go -destination=mock_interface.go -package=service

type User interface {
	Create(*models.UserRequest) (*models.User, error)
	Get() []models.User
	GetByID(int) (*models.User, error)
	Update(*models.UserUpdateRequest) (*models.User, error)
	Delete(int) error

	Search(filters *models.Filters) []models.User
}
