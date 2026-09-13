package organization

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MembershipRepository struct {
	db DBTX
}

func NewMembershipRepository(db *pgxpool.Pool) *MembershipRepository {
	return &MembershipRepository{db: db}
}

var ErrMembershipNotFound = errors.New("membership not found")

func (r *MembershipRepository) Create(ctx context.Context, organizationID, userID uuid.UUID) (*Membership, error) {
	var membership Membership
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO organization-memberships
		(organization_id, user_id)
		VALUES ($1, $2)
		RETURNING id, organization_id, user_id, status, created_at, updated_at`,
		organizationID,
		userID,
	).Scan(
		&membership.ID,
		&membership.OrganizationID,
		&membership.UserID,
		&membership.Status,
		&membership.CreatedAt,
		&membership.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &membership, nil
}
