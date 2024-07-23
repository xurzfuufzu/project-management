package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management-service/internal/domain"
)

type TasksRepo struct {
	db *pgxpool.Pool
}

func NewTasksRepo(db *pgxpool.Pool) *TasksRepo {
	return &TasksRepo{
		db: db,
	}
}

func (t TasksRepo) GetAll(ctx context.Context) ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TasksRepo) Create(ctx context.Context, task domain.Task) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (t TasksRepo) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TasksRepo) Update(ctx context.Context, task domain.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TasksRepo) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (t TasksRepo) SearchByTitle(ctx context.Context, title string) ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TasksRepo) SearchByStatus(ctx context.Context, status domain.Status) ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TasksRepo) SearchByPriority(ctx context.Context, priority domain.Priority) ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TasksRepo) SearchByAssigneeID(ctx context.Context, assigneeID string) ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TasksRepo) SearchByProjectID(ctx context.Context, projectID string) ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}
