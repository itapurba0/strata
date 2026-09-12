package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

var ErrEmailExists = errors.New("user with this email already exists")
var ErrNotFound = errors.New("User not found")

func (r *Repository) Create(ctx context.Context, user *User) (*User, error) {
	var createdUser User

	err := r.db.QueryRow(
		ctx,
		`
		INSERT INTO users (id, email, name, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, name, password_hash, created_at, updated_at
		`,
		user.ID,
		user.Email,
		user.Name,
		user.PasswordHash,
	).Scan(
		&createdUser.ID,
		&createdUser.Email,
		&createdUser.Name,
		&createdUser.PasswordHash,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, ErrEmailExists
		}
		return nil, err
	}

	return &createdUser, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	var user User

	err := r.db.QueryRow(
		ctx,
		`SELECT id, email, name, created_at, updated_at
		FROM users
		where id = $1
		`,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	
	err := r.db.QueryRow(
		ctx,
		`SELECT id, email, name, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil

}
