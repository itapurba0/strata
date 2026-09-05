package organization

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Organization struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}
var ErrNotFound = errors.New("organization not found")

func (r *Repository) Create(ctx context.Context, name string) (*Organization, error) {
	id := uuid.New()

	var organization Organization

	err := r.db.QueryRow(
		ctx,
		"INSERT INTO organizations (id, name) VALUES ($1, $2) RETURNING id, name, created_at, updated_at",
		id,
		name,
	).Scan(
		&organization.ID,
		&organization.Name,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &organization, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*Organization, error) {
	var organization Organization

	err := r.db.QueryRow(
		ctx,
		`
		SELECT id, name, created_at, updated_at
		FROM organizations
		WHERE id = $1
		`,
		id,
	).Scan(
		&organization.ID,
		&organization.Name,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &organization, nil
}