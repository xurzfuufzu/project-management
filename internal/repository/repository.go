package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management-service/internal/domain"
	"project-management-service/internal/repository/postgresql"
)

type Users interface {
	Create(ctx context.Context, user *domain.User) (string, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, user *domain.User) error
	SearchByName(ctx context.Context, name string) ([]domain.User, error)
	SearchByEmail(ctx context.Context, email string) (*domain.User, error)
	GetProjectsByUserID(ctx context.Context, id string) ([]domain.Project, error)
}

type Projects interface {
	GetAll(ctx context.Context) ([]domain.Project, error)
	Create(ctx context.Context, project *domain.Project) (string, error)
	GetByID(ctx context.Context, id string) (*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id string) error
	SearchByTitle(ctx context.Context, title string) (*domain.Project, error)
	SearchByManagerID(ctx context.Context, managerID string) ([]domain.Project, error)
}

type Repositories struct {
	Users
	Projects
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Users:    postgresql.NewUsersRepo(db),
		Projects: postgresql.NewProjectRepo(db),
	}
}
