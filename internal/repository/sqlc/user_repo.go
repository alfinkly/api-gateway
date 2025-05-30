package sqlc

import (
	"context"
	"github.com/alfinkly/api-gateway/internal/domain"
	"github.com/alfinkly/api-gateway/internal/repository/sqlc/db"
)

type UserRepo struct {
	q *db.Queries
}

func (r *UserRepo) SelectById(ctx context.Context, id int32) (domain.User, error) {
	user, err := r.q.SelectById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
	}, nil
}
