package organization

import (
	"time"

	"github.com/google/uuid"
)

type  MembershipStatus string
const (
	MembershipPending MembershipStatus = "pending"
	MembershipActive MembershipStatus = "active"
	MembershipRevoked MembershipStatus = "revoked"
)

type Membership struct{
	ID uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	UserID uuid.UUID `json:"user_id"`
	Status MembershipStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}