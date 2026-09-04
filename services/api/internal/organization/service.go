package organization

import (
	"context"

	"fmt"
	"strings"
)

type Service struct{
	repository *Repository
}

func NewService(repository * Repository) *Service{
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, name string) (*Organization, error){
	name = strings.TrimSpace(name)
	if name == ""{
		return nil, fmt.Errorf("organization name cannot be empty")
	}
	
	if len(name) > 100{
		return nil, fmt.Errorf("organization name cannot exceed 100 characters")
	}


	return s.repository.Create(ctx,name)
}