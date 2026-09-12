package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct{
	jwt.RegisteredClaims
}



func NewClaims(userID uuid.UUID, expiration time.Duration) *Claims {
	now := time.Now()
	return &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
}
const  jwtExpiration = time.Hour*24*7
func GenerateJWT(userID uuid.UUID, secretKey string) (string, error) {
         claims := NewClaims(userID, jwtExpiration)
		 jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		 return jwtToken.SignedString([]byte(secretKey))
}