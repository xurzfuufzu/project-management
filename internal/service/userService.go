package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"project-management-service/internal/domain"
	"project-management-service/internal/repository"
	"time"
)

var ErrUserNotFound = errors.New("user not found")

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, input UserInput) (string, error) {
	user := domain.User{
		ID:               uuid.NewString(),
		Name:             input.Name,
		Email:            input.Email,
		RegistrationDate: time.Now(),
		Role:             input.Role,
	}

	return s.repo.Create(ctx, &user)
}

func (s *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) GetByID(ctx context.Context, UserID string) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) Delete(ctx context.Context, UserID string) error {
	return s.repo.Delete(ctx, UserID)
}

func (s *UserService) Update(ctx context.Context, UserID string, input UserInput) error {
	user := domain.User{
		Name:  input.Name,
		Email: input.Email,
		Role:  input.Role,
	}
	return s.repo.Update(ctx, UserID, &user)
}

func (s *UserService) GetProjectsByUserID(ctx context.Context, id string) ([]domain.Project, error) {
	return s.repo.GetProjectsByUserID(ctx, id)
}

func (s *UserService) SearchByName(ctx context.Context, name string) ([]domain.User, error) {
	return s.repo.SearchByName(ctx, name)
}

func (s *UserService) SearchByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.SearchByEmail(ctx, email)
}
