package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/viniciusith/expense-tracker/internal/domain/dto"
	"github.com/viniciusith/expense-tracker/internal/domain/entities"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *dto.CreateUserRequest) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(ctx, query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *entities.User) error {
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, user *entities.User) error {
	query := "DELETE FROM users WHERE id = $1"

	_, err := r.db.Exec(ctx, query, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindById(ctx context.Context, id string) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	row := r.db.QueryRow(ctx, query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := "SELECT * FROM users WHERE email = $1"

	row := r.db.QueryRow(ctx, query, email)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
