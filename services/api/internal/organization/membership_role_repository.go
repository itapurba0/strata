package organization

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type MembershipRoleRepository struct {
	db DBTX
}

func NewMembershipRoleRepository(db DBTX) *MembershipRoleRepository {
	return &MembershipRoleRepository{
		db: db,
	}
}

func (r *MembershipRoleRepository) Assign(
	ctx context.Context,
	membershipID uuid.UUID,
	roleID uuid.UUID,
) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO membership_roles (membership_id, role_id)
		 VALUES ($1, $2)`,
		membershipID,
		roleID,
	)
	if err != nil {
		return fmt.Errorf("failed to assign role to membership: %w", err)
	}

	return nil
}
