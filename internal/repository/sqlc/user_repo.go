package sqlc

import (
	"context"
	"github.com/alfinkly/api-gateway/internal/domain"
	"github.com/alfinkly/api-gateway/internal/repository/sqlc/db"
	"github.com/ulule/deepcopier"
)

type UserRepo struct {
	q *db.Queries
}

func (r *UserRepo) SelectById(ctx context.Context, id int32) (out domain.User, err error) {
	row, err := r.q.SelectUserById(ctx, id)
	if err != nil {
		return
	}
	err = deepcopier.Copy(row).To(&out)
	return
}

func (r *UserRepo) SelectByEmail(ctx context.Context, email string) (out domain.User, err error) {
	row, err := r.q.SelectUserByEmail(ctx, email)
	if err != nil {
		return
	}
	err = deepcopier.Copy(row).To(&out)
	return
}

func (r *UserRepo) CreateUser(ctx context.Context, in domain.User) (out domain.User, err error) {
	var params db.CreateUserParams
	err = deepcopier.Copy(in).To(&params)
	if err != nil {
		return
	}

	dbUser, err := r.q.CreateUser(ctx, params)
	if err != nil {
		return
	}

	err = deepcopier.Copy(dbUser).To(&out)
	return
}

func (r *UserRepo) UpdateUser(ctx context.Context, in domain.User) (out domain.User, err error) {
	var params db.UpdateUserParams
	err = deepcopier.Copy(in).To(&params)
	if err != nil {
		return
	}

	dbUser, err := r.q.UpdateUser(ctx, params)
	if err != nil {
		return
	}

	err = deepcopier.Copy(dbUser).To(&out)
	return
}

func (r *UserRepo) DeleteById(ctx context.Context, id int32) error {
	return r.q.DeleteUserById(ctx, id)
}
