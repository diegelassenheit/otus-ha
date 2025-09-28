package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"social_network/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, email, passwordHash string) (*domain.User, error)
	CreateProfile(ctx context.Context, firstName, secondName string, birthdate time.Time, biography, city, passwordHash string) (*domain.User, error)
	GetById(ctx context.Context, id string) (*domain.User, error)
}

type userRepo struct {
	pool *pgxpool.Pool
}

func (r *userRepo) Create(ctx context.Context, email, passwordHash string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var u domain.User
	err := r.pool.QueryRow(ctx, `
        INSERT INTO users (email, password_hash, first_name, second_name, birthdate, biography, city)
        VALUES ($1, $2, '', '', NOW(), '', '')
        RETURNING id, email, password_hash, created_at`, email, passwordHash).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // unique_violation
			return nil, ErrEmailAlreadyExists
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) GetById(ctx context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var u domain.User
	err := r.pool.QueryRow(ctx, `
        SELECT id, first_name, second_name, birthdate, biography, city, password_hash, created_at
        FROM users
        WHERE id = $1
    `, id).Scan(&u.ID, &u.FirstName, &u.SecondName, &u.Birthdate, &u.Biography, &u.City, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) CreateProfile(ctx context.Context, firstName, secondName string, birthdate time.Time, biography, city, passwordHash string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	id := uuid.New().String()
	var u domain.User
	err := r.pool.QueryRow(ctx, `
        INSERT INTO users (id, password_hash, first_name, second_name, birthdate, biography, city)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, first_name, second_name, birthdate, biography, city, created_at
    `, id, passwordHash, firstName, secondName, birthdate, biography, city).Scan(&u.ID, &u.FirstName, &u.SecondName, &u.Birthdate, &u.Biography, &u.City, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepo{pool: pool}
}

var ErrEmailAlreadyExists = errors.New("email already exists")
