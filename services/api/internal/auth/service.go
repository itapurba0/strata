package auth

import (
	"context"
	"errors"

	"github.com/itapurba0/strata/services/api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepository *user.Repository
	secretKey	  string
}
type LoggedUser struct {
	User  *user.User
	Token string
}

func NewService(userRepository *user.Repository, secretKey string) *Service {
	return &Service{
		userRepository: userRepository,
		secretKey:      secretKey,
	}
}

type LoginInput struct {
	Email    string
	Password string
}

var ErrInvalidCredentials = errors.New("invalid email or password")

func (s *Service) Login(ctx context.Context, input LoginInput) (*LoggedUser, error) {

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
	token, err := GenerateJWT(foundUser.ID, s.secretKey)
	if err != nil {
		return nil, err
	}

	return &LoggedUser{
		User:  foundUser,
		Token: token,
	}, nil

}
