package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management-service/internal/domain"
	"project-management-service/internal/repository/postgresql"
)

type Users interface {
	Create(ctx context.Context, user domain.User) (string, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetALL(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id string) error
	GetTasksByUserID(ctx context.Context, id string) ([]domain.Task, error)
	SearchByName(ctx context.Context, name string) ([]domain.User, error)
	SearchByEmail(ctx context.Context, email string) ([]domain.User, error)
}

type Tasks interface {
	GetAll(ctx context.Context) ([]domain.Task, error)
	Create(ctx context.Context, task domain.Task) (string, error)
	GetByID(ctx context.Context, id string) (*domain.Task, error)
	Update(ctx context.Context, task domain.Task) error
	Delete(ctx context.Context, id string) error
	SearchByTitle(ctx context.Context, title string) ([]domain.Task, error)
	SearchByStatus(ctx context.Context, status domain.Status) ([]domain.Task, error)
	SearchByPriority(ctx context.Context, priority domain.Priority) ([]domain.Task, error)
	SearchByAssigneeID(ctx context.Context, assigneeID string) ([]domain.Task, error)
	SearchByProjectID(ctx context.Context, projectID string) ([]domain.Task, error)
}

type Projects interface {
	GetAll(ctx context.Context) ([]domain.Project, error)
	Create(ctx context.Context, project domain.Project) (string, error)
	GetByID(ctx context.Context, id string) (*domain.Project, error)
	Update(ctx context.Context, project domain.Project) error
	Delete(ctx context.Context, id string) error
	GetTasksByProjectID(ctx context.Context, projectID string) ([]domain.Task, error)
	SearchByTitle(ctx context.Context, title string) ([]domain.Project, error)
	SearchByManagerID(ctx context.Context, managerID string) ([]domain.Project, error)
}

type Repositories struct {
	Users    Users
	Tasks    Tasks
	Projects Projects
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Users:    postgresql.NewUsersRepo(db),
		Tasks:    postgresql.NewTasksRepo(db),
		Projects: postgresql.NewProjectRepo(db),
	}
}
