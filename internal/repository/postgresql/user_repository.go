package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management-service/internal/domain"
)

type UsersRepo struct {
	db *pgxpool.Pool
}

func NewUsersRepo(db *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Create(ctx context.Context, u domain.User) (string, error) {
	q := `INSERT INTO public.users (name, email, registration_date, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id string
	err := r.db.QueryRow(ctx, q, u.Name, u.Email, u.RegistrationDate, u.Role).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert user: %w", err)
	}
	return id, nil
}

func (u UsersRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepo) GetALL(ctx context.Context) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepo) Update(ctx context.Context, user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepo) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepo) GetTasksByUserID(ctx context.Context, id string) ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepo) SearchByName(ctx context.Context, name string) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepo) SearchByEmail(ctx context.Context, email string) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}
