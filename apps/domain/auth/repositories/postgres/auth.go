package postgres

import (
	"backend/apps/domain/auth/models"
	"backend/apps/domain/auth/repositories"
	"context"
	"database/sql"
)

var (
	queryCreate = `
		INSERT INTO auth (email, password)
		VALUES ($1, $2)
	`

	queryFindByEmail = `
		SELECT id, email, password
		FROM auth
		WHERE email=$1
	`
)

type authRepo struct {
	db *sql.DB
}

// Create implements repositories.AuthRepo
func (a *authRepo) Create(ctx context.Context, auth *models.Auth) error {
	stmt, err := a.db.Prepare(queryCreate)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(auth.Email, auth.Password)
	if err != nil {
		return err
	}

	return nil

}

// FindByEmail implements repositories.AuthRepo
func (a *authRepo) FindByEmail(ctx context.Context, email string) (*models.Auth, error) {
	stmt, err := a.db.Prepare(queryFindByEmail)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(email)

	var auth = new(models.Auth)
	err = row.Scan(&auth.Id, &auth.Email, &auth.Password)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func NewAuthRepo(db *sql.DB) repositories.AuthRepo {
	return &authRepo{
		db: db,
	}
}
