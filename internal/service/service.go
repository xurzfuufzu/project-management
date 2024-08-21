package service

import (
	"context"
	"project-management-service/internal/domain"
	"project-management-service/internal/repository"
	"time"
)

type UserInput struct {
	Name  string
	Email string
	Role  string
}

type Users interface {
	Create(ctx context.Context, input UserInput) (string, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, UserID string) (*domain.User, error)
	Delete(ctx context.Context, UserID string) error
	Update(ctx context.Context, UserID string, input UserInput) error
	SearchByName(ctx context.Context, name string) ([]domain.User, error)
	SearchByEmail(ctx context.Context, email string) (*domain.User, error)
}

type ProjectInput struct {
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	ManagerID   string
}

type Projects interface {
	GetAll(ctx context.Context) ([]domain.Project, error)
	Create(ctx context.Context, input ProjectInput) (string, error)
	GetByID(ctx context.Context, projectID string) (*domain.Project, error)
	Update(ctx context.Context, input ProjectInput) error
	Delete(ctx context.Context, projectID string) error
	SearchByTitle(ctx context.Context, title string) (*domain.Project, error)
	SearchByManagerID(ctx context.Context, managerID string) ([]domain.Project, error)
}

type Services struct {
	Users
	Projects
}

type ServicesDependencies struct {
	Repos *repository.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Users:    NewUserService(deps.Repos.Users),
		Projects: NewProjectService(deps.Repos.Projects),
	}
}
