package user

import (
	"context"
	"github.com/alfinkly/api-gateway/internal/domain"
	"github.com/alfinkly/api-gateway/internal/port"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo port.UserRepository
}

func NewService(repo port.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Register(ctx context.Context, resp domain.RegisterResponse) (user domain.User, err error) {
	passBytes, err := bcrypt.GenerateFromPassword([]byte(resp.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	u := domain.User{Username: resp.Username, Name: resp.Name, Surname: resp.Surname,
		Email: resp.Email, Password: string(passBytes)}
	return s.repo.CreateUser(ctx, u)
}
