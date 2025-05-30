package user

import (
	"context"
	"github.com/alfinkly/api-gateway/internal/domain"
	"github.com/alfinkly/api-gateway/internal/port"
)

type Service struct {
	repo port.UserRepository
}

func NewService(repo port.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) Register(ctx context.Context, name, email, passwordHash string) (domain.User, error) {
	u := domain.User{Name: name, Email: email, Password: passwordHash}
	return s.repo.Create(ctx, u)
}
