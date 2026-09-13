package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/itapurba0/strata/services/api/internal/auth"
	"github.com/itapurba0/strata/services/api/internal/middleware"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	secretKey := "test-secret-key"

	userID := uuid.New()

	token, err := auth.GenerateJWT(userID, secretKey)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	authMiddleware := middleware.NewAuthMiddleware(secretKey)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUserID, ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			t.Fatal("user ID not found in context")
		}

		if gotUserID != userID {
			t.Fatalf("expected user ID %s, got %s", userID, gotUserID)
		}

		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/users/me",
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	handler := authMiddleware.Middleware(nextHandler)

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	secretKey := "test-secret-key"

	authMiddleware := middleware.NewAuthMiddleware(secretKey)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next handler should not be called for invalid token")
	})	

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/users/me",
		nil,
	)
	
	req.Header.Set("Authorization", "Bearer invalid-token")

	rec := httptest.NewRecorder()

	handler := authMiddleware.Middleware(nextHandler)

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rec.Code)
	}
}

func TestAuthMiddleware_WrongSecret(t *testing.T) {
	secretKey := "test-secret-key"

	userID := uuid.New()

	token, err := auth.GenerateJWT(userID, secretKey)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	wrongSecretKey := "wrong-secret-key"
	authMiddleware := middleware.NewAuthMiddleware(wrongSecretKey)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next handler should not be called for wrong secret key")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/users/me",
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	handler := authMiddleware.Middleware(nextHandler)

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rec.Code)
	}
}