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

func NewPostgreRepository(pool *pgxpool.Pool) *PostgreRepository {
	return &PostgreRepository{
		pool: pool,
	}
}

func (r *PostgreRepository) Create(ctx context.Context, newClass domains.Class) (domains.Class, error) {

	query := `
		INSERT INTO classes (name,join_code) VALUES ($1, $2)
		RETURNING
			id,
			name,
			join_code,
			created_at,
			updated_at
	`

	var model classModel

	err := r.pool.QueryRow(ctx, query, newClass.Name, newClass.JoinCode).Scan(
		&model.ID,
		&model.Name,
		&model.JoinCode,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domains.Class{}, apperrors.ErrJoinCodeAlreadyTaken
		}

		return domains.Class{}, err
	}

	return toDomain(model), nil
}

func (r *PostgreRepository) GetByID(ctx context.Context, id int) (domains.Class, error) {

	query := `
		SELECT id, name, join_code, created_at, updated_at FROM classes
		WHERE id = $1
	`

	var model classModel

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Name,
		&model.JoinCode,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return domains.Class{},
			apperrors.ErrClassNotFound
	}

	if err != nil {
		return domains.Class{}, err
	}

	return toDomain(model), nil
}

func (r *PostgreRepository) GetByJoinCode(ctx context.Context, joinCode string) (domains.Class, error) {
	query := `
		SELECT id, name, join_code, created_at, updated_at FROM classes
		WHERE join_code = $1
	`

	var model classModel

	err := r.pool.QueryRow(ctx, query, joinCode).Scan(
		&model.ID,
		&model.Name,
		&model.JoinCode,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return domains.Class{},
			apperrors.ErrClassNotFound
	}

	if err != nil {
		return domains.Class{}, err
	}

	return toDomain(model), nil
}

func (r *PostgreRepository) Update(ctx context.Context, classChanged domains.Class) (domains.Class, error) {
	query := `
		UPDATE classes 
		SET name = $1, join_code = $2, updated_at = NOW() 
		WHERE id = $3
		RETURNING id, name, join_code, created_at, updated_at
	`

	var model classModel

	err := r.pool.QueryRow(ctx, query, classChanged.Name, classChanged.JoinCode, classChanged.ID).Scan(
		&model.ID,
		&model.Name,
		&model.JoinCode,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return domains.Class{},
			apperrors.ErrClassNotFound
	}

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domains.Class{},
				apperrors.ErrJoinCodeAlreadyTaken
		}

		return domains.Class{}, err
	}

	return toDomain(model), nil
}
