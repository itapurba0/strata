package organization

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Organization struct {
	ID         uuid.UUID
	Name       string
	CreatedAt  string
	UpdatedAt  string
}

type Repository struct{
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository{
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, name string) (*Organization, error){
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

	if err != nil{
		return nil, err
	}

	return &organization,nil;
}