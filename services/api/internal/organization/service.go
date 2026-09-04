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
