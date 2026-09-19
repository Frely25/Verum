package repository

import (
	"context"
	"testing"
	"time"
)

func TestMemoryRepositoryCreateAndGet(t *testing.T) {
	repo := NewMemoryRepository()

	ctx := context.Background()

	token, expiresAt, err := repo.Create(ctx, 42, time.Hour)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if token == "" {
		t.Fatal("token is empty")
	}

	if expiresAt.Before(time.Now()) {
		t.Fatal("session already expired")
	}

	userID, err := repo.Get(ctx, token)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if userID != 42 {
		t.Errorf("expected user id 42, got %d", userID)
	}
}

func TestMemoryRepositoryDelete(t *testing.T) {
	repo := NewMemoryRepository()

	ctx := context.Background()

	token, _, err := repo.Create(ctx, 42, time.Hour)

	if err != nil {
		t.Fatal(err)
	}

	err = repo.Delete(ctx, token)

	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.Get(ctx, token)

	if err == nil {
		t.Fatal("expected error after deleting session")
	}
}

func TestMemoryRepositoryExpiredSession(t *testing.T) {
	repo := NewMemoryRepository()

	ctx := context.Background()

	token, _, err := repo.Create(ctx, 42, -time.Second)

	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.Get(ctx, token)

	if err == nil {
		t.Fatal("expected expired session error")
	}
}
