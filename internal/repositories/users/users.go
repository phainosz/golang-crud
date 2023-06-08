package users

import (
	"context"

	"github.com/phainosz/golang-crud/internal/models"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user models.User) error
	DeleteUserById(ctx context.Context, id uint64) error
	UpdateUser(ctx context.Context, id uint64, user models.User) error
	FindUserById(ctx context.Context, id uint64) (models.User, error)
}
