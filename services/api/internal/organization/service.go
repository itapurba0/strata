package organization

import (
	"context"

	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
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
