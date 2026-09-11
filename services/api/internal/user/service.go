package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Email    string
	Name     string
	Password string
}

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, input CreateUserInput) (*User, error) {
	input.Email = strings.TrimSpace(input.Email)

	if input.Email == "" {
		return nil, fmt.Errorf("user email cannot be empty")
	}

	if len(input.Email) > 100 {
		return nil, fmt.Errorf("user email cannot exceed 100 characters")
	}

	input.Name = strings.TrimSpace(input.Name)

	if input.Name == "" {
		return nil, fmt.Errorf("user name cannot be empty")
	}

	if len(input.Name) > 100 {
		return nil, fmt.Errorf("user name cannot exceed 100 characters")
	}

	if input.Password == "" {
		return nil, fmt.Errorf("user password cannot be empty")
	}

	user := &User{
		ID:    uuid.New(),
		Email: input.Email,
		Name:  input.Name,
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user.PasswordHash = string(passwordHash)

	return s.repository.Create(ctx, user)
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {

	if id == uuid.Nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	return s.repository.GetByID(ctx, id)
}
