package port

import (
	"context"
	"github.com/alfinkly/api-gateway/internal/domain"
)

type UserRepository interface {
	SelectById(ctx context.Context, id int) (domain.User, error)
	SelectByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, u domain.User) (domain.User, error)
	Update(ctx context.Context, u domain.User) (domain.User, error)
	Delete(ctx context.Context, id int) error
}
