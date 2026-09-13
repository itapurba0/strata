package organization

import (
	"context"

	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	repository         *Repository
	membershipRepo     *MembershipRepository
	membershipRoleRepo *MembershipRoleRepository
	roleRepo           *RoleRepository
	db                 *pgxpool.Pool
}

func NewService(
	repository *Repository,
	membershipRepo *MembershipRepository,
	membershipRoleRepo *MembershipRoleRepository,
	roleRepo *RoleRepository,
	db *pgxpool.Pool,
) *Service {
	return &Service{
		repository:         repository,
		membershipRepo:     membershipRepo,
		membershipRoleRepo: membershipRoleRepo,
		roleRepo:           roleRepo,
		db:                 db,
	}
}

type CreateOrganizationInput struct {
	Name string `json:"name"`
}

func (s *Service) Create(ctx context.Context, name string) (*Organization, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, fmt.Errorf("organization name cannot be empty")
	}

	if len(name) > 100 {
		return nil, fmt.Errorf("organization name cannot exceed 100 characters")
	}

	return s.repository.Create(ctx, name)
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Organization, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("invalid organization ID")
	}

	return s.repository.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context, page, limit int) ([]*Organization, int, error) {

	if page < 1 {
		return nil, 0, fmt.Errorf("page number must be greater than 0")
	}

	if limit < 1 || limit > 100 {
		return nil, 0, fmt.Errorf("limit must be between 1 and 100")
	}

	return s.repository.List(ctx, page, limit)
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, name string) (*Organization, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, fmt.Errorf("organization name cannot be empty")
	}

	if len(name) > 100 {
		return nil, fmt.Errorf("organization name cannot exceed 100 characters")
	}

	return s.repository.Update(ctx, id, name)
}

func (s *Service) CreateOrganization(ctx context.Context, userID uuid.UUID, name string) (*Organization, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	organizationRepo := &Repository{db: tx}
	membershipRepo := &MembershipRepository{db: tx}

	organization, err := organizationRepo.Create(ctx, name)
	if err != nil {
		return nil, err
	}

	_, err = membershipRepo.Create(ctx, organization.ID, userID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return organization, nil
}
