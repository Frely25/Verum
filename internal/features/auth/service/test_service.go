package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Frely25/Verum/internal/core/domains"
	apperrors "github.com/Frely25/Verum/internal/core/errors"
	"github.com/Frely25/Verum/internal/features/auth"
	"golang.org/x/crypto/bcrypt"
)

// Создаем фейковый сервис для user
type fakeUserService struct {
	createdUser domains.User
	user        domains.User

	createErr     error
	getByIDErr    error
	getByLoginErr error
}

func (f *fakeUserService) Create(ctx context.Context, user domains.User) (domains.User, error) {

	if f.createErr != nil {
		return domains.User{}, f.createErr
	}

	f.createdUser = user

	user.ID = 1

	return user, nil
}

func (f *fakeUserService) GetByID(ctx context.Context, id int) (domains.User, error) {

	if f.getByIDErr != nil {
		return domains.User{}, f.getByIDErr
	}

	return f.user, nil
}

func (f *fakeUserService) GetByLogin(ctx context.Context, login string) (domains.User, error) {

	if f.getByLoginErr != nil {
		return domains.User{}, f.getByLoginErr
	}

	return f.user, nil
}

// Созад
type fakeSessionRepository struct {
	userID    int
	token     string
	expiresAt time.Time

	createErr error
	getErr    error
	deleteErr error

	deletedToken string
}

func (f *fakeSessionRepository) Create(ctx context.Context, userID int, ttl time.Duration) (string, time.Time, error) {

	if f.createErr != nil {
		return "", time.Time{}, f.createErr
	}

	f.userID = userID

	return f.token, f.expiresAt, nil
}

func (f *fakeSessionRepository) Get(ctx context.Context, token string) (int, error) {

	if f.getErr != nil {
		return 0, f.getErr
	}

	return f.userID, nil
}

func (f *fakeSessionRepository) Delete(ctx context.Context, token string) error {

	if f.deleteErr != nil {
		return f.deleteErr
	}

	f.deletedToken = token

	return nil
}

// ================== TESTS ==================

func TestRegisterSuccess(t *testing.T) {
	users := &fakeUserService{}
	sessions := &fakeSessionRepository{}

	service := NewService(users, sessions, 24*time.Hour)

	input := auth.RegisterInput{
		Login:       "frely",
		Password:    "12345678",
		DisplayName: "Frely",
	}

	createdUser, err := service.Register(context.Background(), input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if createdUser.Login != "frely" {
		t.Errorf("expected login frely, got %s", createdUser.Login)
	}

	if createdUser.DisplayName != "Frely" {
		t.Errorf("expected display name Frely, got %s", createdUser.DisplayName)
	}

	if createdUser.PasswordHash == "" {
		t.Fatal("password hash is empty")
	}
	err = bcrypt.CompareHashAndPassword([]byte(createdUser.PasswordHash), []byte(input.Password))

	if err != nil {
		t.Fatal("password hash does not match password")
	}

	if createdUser.PasswordHash == input.Password {
		t.Fatal("plain password was stored instead of hash")
	}
}

func TestRegisterInvalidInput(t *testing.T) {
	users := &fakeUserService{}
	sessions := &fakeSessionRepository{}

	service := NewService(
		users,
		sessions,
		24*time.Hour,
	)

	_, err := service.Register(context.Background(),
		auth.RegisterInput{
			Login:       "",
			Password:    "12345678",
			DisplayName: "Frely",
		},
	)

	if !errors.Is(err, apperrors.ErrInvalidAuthInput) {
		t.Fatalf("expected ErrInvalidAuthInput, got %v", err)
	}
}

func TestLoginSuccess(t *testing.T) {
	password := "12345678"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		t.Fatal(err)
	}

	users := &fakeUserService{
		user: domains.User{
			ID:           10,
			Login:        "frely",
			PasswordHash: string(hash),
			DisplayName:  "Frely",
		},
	}

	expiresAt := time.Now().Add(24 * time.Hour)

	sessions := &fakeSessionRepository{
		token:     "test-token",
		expiresAt: expiresAt,
	}

	service := NewService(
		users,
		sessions,
		24*time.Hour,
	)

	result, err := service.Login(context.Background(),
		auth.LoginInput{
			Login:    "frely",
			Password: password,
		},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.User.ID != 10 {
		t.Errorf("expected user id 10, got %d", result.User.ID)
	}

	if result.Token != "test-token" {
		t.Errorf("expected test-token, got %s", result.Token)
	}

	if sessions.userID != 10 {
		t.Errorf("expected session user id 10, got %d", sessions.userID)
	}
}

func TestLoginWrongPassword(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("correct-password"), bcrypt.DefaultCost)

	if err != nil {
		t.Fatal(err)
	}

	users := &fakeUserService{
		user: domains.User{
			ID:           1,
			Login:        "frely",
			PasswordHash: string(hash),
		},
	}

	sessions := &fakeSessionRepository{}

	service := NewService(
		users,
		sessions,
		24*time.Hour,
	)

	_, err = service.Login(context.Background(),
		auth.LoginInput{
			Login:    "frely",
			Password: "wrong-password",
		},
	)

	if !errors.Is(err, apperrors.ErrInvalidCredentials) {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}

func TestLoginUserNotFound(t *testing.T) {
	users := &fakeUserService{getByLoginErr: apperrors.ErrUserNotFound}
	sessions := &fakeSessionRepository{}

	service := NewService(
		users,
		sessions,
		24*time.Hour,
	)

	_, err := service.Login(context.Background(),
		auth.LoginInput{
			Login:    "unknown",
			Password: "12345678",
		},
	)

	if !errors.Is(err, apperrors.ErrInvalidCredentials) {
		t.Fatalf(
			"expected ErrInvalidCredentials, got %v",
			err,
		)
	}
}

func TestAuthenticateSuccess(t *testing.T) {
	users := &fakeUserService{
		user: domains.User{
			ID:          5,
			Login:       "frely",
			DisplayName: "Frely",
		},
	}

	sessions := &fakeSessionRepository{userID: 5}

	service := NewService(
		users,
		sessions,
		24*time.Hour,
	)

	user, err := service.Authenticate(context.Background(), "valid-token")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.ID != 5 {
		t.Errorf("expected user id 5, got %d", user.ID)
	}
}

func TestAuthenticateInvalidSession(t *testing.T) {
	users := &fakeUserService{}

	sessions := &fakeSessionRepository{getErr: apperrors.ErrSessionNotFound}

	service := NewService(
		users,
		sessions,
		24*time.Hour,
	)

	_, err := service.Authenticate(context.Background(), "invalid-token")

	if !errors.Is(err, apperrors.ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestLogout(t *testing.T) {
	users := &fakeUserService{}
	sessions := &fakeSessionRepository{}

	service := NewService(
		users,
		sessions,
		24*time.Hour,
	)

	err := service.Logout(context.Background(), "token-to-delete")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if sessions.deletedToken != "token-to-delete" {
		t.Errorf("expected token-to-delete, got %s", sessions.deletedToken)
	}
}
