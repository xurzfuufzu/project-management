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
	"strconv"
)

var ErrUserNotFound = errors.New("user not found")

type UsersRepo struct {
	db *pgxpool.Pool
}

func NewUsersRepo(db *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Create(ctx context.Context, u *domain.User) (string, error) {
	q := `INSERT INTO public.users (id, name, email, registration_date, role)
		  VALUES ($1, $2, $3, $4, $5)
          RETURNING id`

	var id string

	err := r.db.QueryRow(ctx, q, u.ID, u.Name, u.Email, u.RegistrationDate, u.Role).Scan(&id)
	if err != nil {
		log.Debug("err: %v", err)
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return "0", repoerrs.ErrAlreadyExists
			}
		}
	}

	return id, nil
}

func (r *UsersRepo) GetAll(ctx context.Context) ([]domain.User, error) {
	q := `SELECT id, name, email, registration_date, role FROM public.users`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.RegistrationDate, &u.Role); err != nil {
			log.Printf("%s", err)
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UsersRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	q := `SELECT id, name, email, registration_date, role FROM public.users WHERE id=$1`

	var u domain.User
	err := r.db.QueryRow(ctx, q, id).Scan(&u.ID, &u.Name, &u.Email, &u.RegistrationDate, &u.Role)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, repoerrs.ErrNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (r *UsersRepo) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM public.users WHERE id=$1`

	cmdTag, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (r *UsersRepo) Update(ctx context.Context, id string, user *domain.User) error {
	userID, err := strconv.Atoi(id)

	q := `UPDATE public.users
	      SET name=$1, email=$2,role=$3
	      WHERE id=$4`

	_, err = r.db.Exec(ctx, q, user.Name, user.Email, user.Role, userID)
	return err
}

func (r *UsersRepo) GetProjectsByUserID(ctx context.Context, id string) ([]domain.Project, error) {
	q := `SELECT
   		id, title, description, start_date, end_date
	      	FROM public.projects WHERE manager_id = $1`

	rows, err := r.db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []domain.Project
	for rows.Next() {
		var p domain.Project
		_ = rows.Scan(
			&p.ID, &p.Title, &p.Description, &p.StartDate, &p.EndDate, &p.ManagerID,
		)
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *UsersRepo) SearchByName(ctx context.Context, name string) ([]domain.User, error) {
	q := `SELECT id, name, email, registration_date,role from public.users WHERE name = $1`

	rows, err := r.db.Query(ctx, q, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var u domain.User
		_ = rows.Scan(&u.ID, &u.Name, &u.Email, &u.RegistrationDate, &u.Role)
		users = append(users, u)
	}

	return users, nil
}

func (r *UsersRepo) SearchByEmail(ctx context.Context, email string) (*domain.User, error) {
	q := `SELECT id, name, email, registration_date, role FROM public.users WHERE email = $1`

	row := r.db.QueryRow(ctx, q, email)

	var u domain.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.RegistrationDate, &u.Role); err != nil {
		return nil, err
	}

	return &u, nil
}
