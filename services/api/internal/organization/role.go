package organization

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID *uuid.UUID `json:"organization_id"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
