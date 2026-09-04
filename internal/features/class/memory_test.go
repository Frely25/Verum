package class

import (
	"errors"
	"testing"
)

// Через *testing.T мы можем сказать Go, что тест провалился
// t.Fatal(...)
// t.Fatalf(...)
// t.Error(...)
// t.Errorf(...)

func TestMemoryRepositoryCreate(t *testing.T) {

	// Для первого теста нам нужно:
	// Создать Repository -> Создать Class{Name: "Go"} -> repo.Create()
	// Хотим получить:
	// ID = 1
	// Name = Go

	repo := NewMemoryRepository()
	input := Class{
		Name: "Go Backend",
	}

	created, err := repo.Create(input)
	if err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	if created.ID != 1 {
		t.Errorf("expected ID 1, got %d", created.ID)
	}

	if created.Name != "Go Backend" {
		t.Errorf(
			"expected name %q, got %q",
			"Go Backend",
			created.Name,
		)
	}
}

func TestMemoryRepositoryCreateAssignsDifferentIDs(t *testing.T) {
	repo := NewMemoryRepository()

	first, err := repo.Create(Class{
		Name: "Go",
	})
	if err != nil {
		t.Fatal(err)
	}

	second, err := repo.Create(Class{
		Name: "Python",
	})
	if err != nil {
		t.Fatal(err)
	}

	if first.ID == second.ID {
		t.Errorf(
			"expected different IDs, got %d and %d",
			first.ID,
			second.ID,
		)
	}
}

func TestMemoryRepositoryGetAll(t *testing.T) {
	repo := NewMemoryRepository()

	names_for_classes := []string{"Go", "Python", "Java"}

	// Добавил в массив несколько классов
	for _, value := range names_for_classes {
		_, err := repo.Create(Class{
			Name: value,
		})

		if err != nil {
			t.Fatal(err)
		}
	}

	classes, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() returned error: %v", err)
	}

	if len(classes) != 3 {
		t.Fatalf(
			"expected 3 classes, got %d",
			len(classes),
		)
	}

	for index, value := range names_for_classes {
		if classes[index].Name != value {
			t.Fatalf("expected classes name {%q}, got {%q}", value, classes[index].Name)
		}
	}
}

func TestMemoryRepositoryGetByID(t *testing.T) {
	repo := NewMemoryRepository()

	created, err := repo.Create(Class{
		Name: "Go Backend",
	})
	if err != nil {
		t.Fatal(err)
	}

	found, err := repo.GetByID(created.ID)
	if err != nil {
		t.Fatal(err)
	}

	if found != created {
		t.Errorf(
			"expected %+v, got %+v",
			created,
			found,
		)
	}
}

func TestMemoryRepositoryGetByIDNotFound(t *testing.T) {
	repo := NewMemoryRepository()

	_, err := repo.GetByID(-123)

	if !errors.Is(err, ErrClassNotFound) {
		t.Errorf(
			"expected ErrClassNotFound, got %v",
			err,
		)
	}
}
