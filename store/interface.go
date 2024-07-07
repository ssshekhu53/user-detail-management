package store

import "github.com/ssshekhu53/user-detail-management/models"

type User interface {
	Create(*models.UserRequest) int64
	Get(filters *models.Filters) []models.User
	GetByID(int64) (*models.User, error)
	Update(*models.User)
	Delete(int64)
}
