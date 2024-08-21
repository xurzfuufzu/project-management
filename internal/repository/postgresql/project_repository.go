package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"project-management-service/internal/domain"
	"project-management-service/internal/repository/repoerrs"
)

type ProjectsRepo struct {
	db *pgxpool.Pool
}

func NewProjectRepo(db *pgxpool.Pool) *ProjectsRepo {
	return &ProjectsRepo{
		db: db,
	}
}

func (r *ProjectsRepo) GetAll(ctx context.Context) ([]domain.Project, error) {
	q := `SELECT id, title, description, start_date, end_date, manager_id FROM public.projects`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []domain.Project

	for rows.Next() {
		var p domain.Project
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.StartDate, &p.EndDate, &p.ManagerID); err != nil {
			log.Printf("%s", err)
			return nil, err
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectsRepo) Create(ctx context.Context, project *domain.Project) (string, error) {
	q := `INSERT INTO public.projects (id, name, description, start_date, end_date, manager_id)
	      VALUES ($1, $2, $3, $4, $5, $6)
		  RETURNING id`

	var id string

	err := r.db.QueryRow(ctx, q,
		project.ID, project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID).
		Scan(&id)
	if err != nil {
		log.Debugf("err: %v", err)
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return "0", repoerrs.ErrAlreadyExists
			}
		}
	}

	return id, nil
}

func (r *ProjectsRepo) GetByID(ctx context.Context, id string) (*domain.Project, error) {
	q := `SELECT id, title, description, start_date, end_date, manager_id FROM public.projects WHERE id = $1`

	row := r.db.QueryRow(ctx, q, id)
	var project domain.Project
	err := row.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, repoerrs.ErrNotFound
		}
		return nil, err
	}

	return &project, nil
}

func (r *ProjectsRepo) Update(ctx context.Context, project *domain.Project) error {
	panic("implement me")
}

func (r *ProjectsRepo) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM public.projects WHERE id=$1`

	cmdTag, err := r.db.Exec(ctx, q, id)

	if err != nil {
		return fmt.Errorf("could not delete project: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return repoerrs.ErrNotFound
	}

	return nil
}

func (r *ProjectsRepo) SearchByTitle(ctx context.Context, title string) (*domain.Project, error) {
	q := `SELECT id, title, description, start_date, end_date, manager_id FROM public.projects WHERE title=$1`

	row := r.db.QueryRow(ctx, q, title)

	var p domain.Project
	if err := row.Scan(&p.ID, &p.Title, &p.Description, &p.StartDate, &p.EndDate, &p.ManagerID); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProjectsRepo) SearchByManagerID(ctx context.Context, managerID string) ([]domain.Project, error) {
	q := `SELECT id, title, description, start_date, end_date, manager_id FROM public.projects WHERE manager_id=$1`

	rows, err := r.db.Query(ctx, q, managerID)
	if err != nil {
		return nil, err
	}

	var projects []domain.Project

	for rows.Next() {
		var p domain.Project
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.StartDate, &p.EndDate, &p.ManagerID); err != nil {
			log.Printf("%s", err)
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}
