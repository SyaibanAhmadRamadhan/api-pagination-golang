package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/internal/http-protocol/exception"
	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/src/modules/profile/entities"
	"github.com/rs/zerolog/log"
)

type ProfileRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, profile entities.Profile) (*entities.Profile, error)
	Update(ctx context.Context, tx *sql.Tx, profile entities.Profile) (*entities.Profile, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindByID(ctx context.Context, db *sql.DB, id string) (*entities.Profile, error)
	FindAll(ctx context.Context, db *sql.DB, limit, offset int, sort, filter string) (*[]entities.Profile, error)
	CountData(ctx context.Context, db *sql.DB) (int, error)
}

type ProfileRepositoryImpl struct{}

func NewProfileRepositoryImpl() ProfileRepository {
	return &ProfileRepositoryImpl{}
}

func (repo *ProfileRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, profile entities.Profile) (*entities.Profile, error) {
	SQL := "SELECT email, id FROM profiles WHERE email = ? LIMIT 1"
	row, err := tx.QueryContext(ctx, SQL, profile.Email)
	if err != nil {
		log.Fatal().Msgf("error query context : %v", err)
		return nil, err
	}
	row.Close()

	if row.Next() {
		return nil, exception.BadRequest(map[string]map[string]string{
			"email": {
				"unique": "email is alvailable",
			},
		})
	}

	SQL = "INSERT INTO profiles(id, name, gender, phone, email, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, profile.ID, profile.Name, profile.Gender, profile.Phone, profile.Email, profile.CreatedAt, profile.UpdatedAt)
	if err != nil {
		log.Fatal().Msgf("error exec context : %v", err)
		return nil, err
	}

	return &profile, nil
}

func (repo *ProfileRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, profile entities.Profile) (*entities.Profile, error) {
	SQL := "SELECT email, id FROM profiles WHERE email = ? AND id <> ? LIMIT 1"
	row, err := tx.QueryContext(ctx, SQL, profile.Email, profile.ID)
	if err != nil {
		log.Err(err).Msgf("error query context : %v", err)
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		return nil, exception.BadRequest(map[string]map[string]string{
			"email": {
				"unique": "email is alvailable",
			},
		})
	}

	SQL = "UPDATE profiles SET name=?, gender=?, phone=?, email=?, updated_at=? WHERE id=?"
	_, err = tx.ExecContext(ctx, SQL, profile.Name, profile.Gender, profile.Phone, profile.Email, profile.UpdatedAt, profile.ID)
	if err != nil {
		log.Err(err).Msgf("error exec context : %v", err)
		return nil, err
	}
	return &profile, nil
}

func (repo *ProfileRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM profiles WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return exception.NotFound("profile not found")
	}
	return nil
}

func (repo *ProfileRepositoryImpl) FindByID(ctx context.Context, db *sql.DB, id string) (*entities.Profile, error) {
	SQL := "SELECT id, name, gender, phone, email, created_at, updated_at FROM profiles WHERE id = ?"
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		log.Fatal().Msgf("error query context : %v", err)
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		var profileEntity entities.Profile
		if err := row.Scan(
			&profileEntity.ID,
			&profileEntity.Name,
			&profileEntity.Gender,
			&profileEntity.Phone,
			&profileEntity.Email,
			&profileEntity.CreatedAt,
			&profileEntity.UpdatedAt,
		); err != nil {
			return nil, err
		}
		return &profileEntity, nil
	}
	return nil, exception.NotFound("profile not found")
}

func (repo *ProfileRepositoryImpl) FindAll(ctx context.Context, db *sql.DB, limit, offset int, sort, filter string) (*[]entities.Profile, error) {
	SQL := fmt.Sprintf("SELECT id, name, gender, phone, email, created_at, updated_at FROM profiles WHERE created_at >= '%s' ORDER BY created_at %s limit ? offset ?", filter+"%", sort)
	rows, err := db.QueryContext(ctx, SQL, limit, offset)
	if err != nil {
		log.Err(err).Msg("cannot query")
		return nil, err
	}
	defer rows.Close()

	var profiles []entities.Profile
	for rows.Next() {
		var profile entities.Profile
		if err := rows.Scan(
			&profile.ID,
			&profile.Name,
			&profile.Gender,
			&profile.Phone,
			&profile.Email,
			&profile.CreatedAt,
			&profile.UpdatedAt,
		); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return &profiles, nil
}

func (repo *ProfileRepositoryImpl) CountData(ctx context.Context, db *sql.DB) (int, error) {
	SQL := "SELECT count(id) FROM profiles"
	row, err := db.QueryContext(ctx, SQL)
	if err != nil {
		return 0, err
	}

	var count int
	for row.Next() {
		if err := row.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}
