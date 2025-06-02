package port

import (
	"context"
	"github.com/alfinkly/api-gateway/internal/domain"
)

type UserRepository interface {
	SelectById(ctx context.Context, id int32) (domain.User, error)
	SelectByEmail(ctx context.Context, email string) (domain.User, error)
	CreateUser(ctx context.Context, in domain.User) (domain.User, error)
	UpdateUser(ctx context.Context, in domain.User) (domain.User, error)
	DeleteUserById(ctx context.Context, id int32) error
}
