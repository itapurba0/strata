package organization

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var ErrRoleNotFound = errors.New("role not found")

type RoleRepository struct {
	db DBTX
}

func NewRoleRepository(db DBTX) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) GetSystemRoleByName(
	ctx context.Context,
	name string,
) (*Role, error) {
	var role Role

	err := r.db.QueryRow(
		ctx,
		`SELECT
			id,
			organization_id,
			name,
			description,
			created_at,
			updated_at
		FROM roles
		WHERE name = $1
		  AND organization_id IS NULL`,
		name,
	).Scan(
		&role.ID,
		&role.OrganizationID,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRoleNotFound
		}

		return nil, fmt.Errorf("failed to get system role: %w", err)
	}

	return &role, nil
}
