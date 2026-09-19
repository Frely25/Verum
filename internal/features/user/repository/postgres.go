package repository

import (
	"context"
	"errors"

	"github.com/Frely25/Verum/internal/core/domains"
	apperrors "github.com/Frely25/Verum/internal/core/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreRepository struct {
	pool *pgxpool.Pool
}

func NewPostgreRepository(db *pgxpool.Pool) *PostgreRepository {
	return &PostgreRepository{
		pool: db,
	}
}

func (r *PostgreRepository) Create(ctx context.Context, newUser domains.User) (domains.User, error) {
	query := `
		INSERT INTO users (login, password_hash, display_name) VALUES ($1, $2, $3)
		RETURNING (id, login, password_hash, display_name, created_at, updated_at)
	`
	var model userModel

	row := r.pool.QueryRow(ctx, query, newUser.Login, newUser.PasswordHash, newUser.DisplayName)

	err := row.Scan(
		&model.ID,
		&model.Login,
		&model.PasswordHash,
		&model.DisplayName,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domains.User{}, apperrors.ErrLoginAlreadyTaken
		}
		return domains.User{}, err
	}

	return model.toDomain(), nil
}

func (r *PostgreRepository) GetByID(ctx context.Context, id int) (domains.User, error) {
	query := `
		SELECT
			id,
			login,
			password_hash,
			display_name,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	var model userModel

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Login,
		&model.PasswordHash,
		&model.DisplayName,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return domains.User{},
			apperrors.ErrUserNotFound
	}

	if err != nil {
		return domains.User{}, err
	}

	return model.toDomain(), nil
}

func (r *PostgreRepository) GetByLogin(ctx context.Context, login string) (domains.User, error) {
	query := `
		SELECT
			id,
			login,
			password_hash,
			display_name,
			created_at,
			updated_at
		FROM users
		WHERE login = $1
	`

	var model userModel

	err := r.pool.QueryRow(ctx, query, login).Scan(
		&model.ID,
		&model.Login,
		&model.PasswordHash,
		&model.DisplayName,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return domains.User{},
			apperrors.ErrUserNotFound
	}

	if err != nil {
		return domains.User{}, err
	}

	return model.toDomain(), nil
}
