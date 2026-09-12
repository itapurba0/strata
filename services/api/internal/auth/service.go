package auth

import (
	"context"
	"errors"

	"github.com/itapurba0/strata/services/api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepository *user.Repository
}

func NewService(userRepository *user.Repository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

type LoginInput struct {
	Email    string
	Password string
}

var ErrInvalidCredentials = errors.New("invalid email or password")

func (s *Service) Login(ctx context.Context, input LoginInput) (*user.User, error) {

	foundUser, err := s.userRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.PasswordHash),
		[]byte(input.Password),
	)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return foundUser, nil

}
