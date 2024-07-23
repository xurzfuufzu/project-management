package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management-service/internal/domain"
)

type ProjectsRepo struct {
	db *pgxpool.Pool
}

func NewProjectRepo(db *pgxpool.Pool) *ProjectsRepo {
	return &ProjectsRepo{
		db: db,
	}
}

func (p ProjectsRepo) GetAll(ctx context.Context) ([]domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectsRepo) Create(ctx context.Context, project domain.Project) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectsRepo) GetByID(ctx context.Context, id string) (*domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectsRepo) Update(ctx context.Context, project domain.Project) error {
	//TODO implement me
	panic("implement me")
}

func (p ProjectsRepo) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProjectsRepo) GetTasksByProjectID(ctx context.Context, projectID string) ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectsRepo) SearchByTitle(ctx context.Context, title string) ([]domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectsRepo) SearchByManagerID(ctx context.Context, managerID string) ([]domain.Project, error) {
	//TODO implement me
	panic("implement me")
}
