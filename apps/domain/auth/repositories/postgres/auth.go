package postgres

import (
	"backend/apps/domain/auth/models"
	"backend/apps/domain/auth/repositories"
	"context"
	"database/sql"
)

var (
	queryCreate = `
		INSERT INTO auth (email, password, img_url)
		VALUES ($1, $2, $3)
	`

	queryFindByEmail = `
		SELECT id, email, password
		FROM auth
		WHERE email=$1
	`

	querySearchByEmail = `
		SELECT DISTINCT(a.email), a.id,COALESCE(a.img_url, ''),f.followingId, f.authId FROM auth as a
		LEFT JOIN follows as f
			ON a.id = f.followingId
		WHERE (a.id <> $1
		OR f.authId=$1)
		AND a.email LIKE $2
	`

	queryFindById = `
		SELECT id, email, COALESCE(img_url, '')
		FROM auth
		WHERE id=$1
	`
)

type authRepo struct {
	db *sql.DB
}

// FindById implements repositories.AuthRepo
func (a *authRepo) FindById(ctx context.Context, authId int) (*models.Auth, error) {
	stmt, err := a.db.Prepare(queryFindById)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(authId)

	var auth = models.Auth{}
	err = row.Scan(
		&auth.Id,
		&auth.Email,
		&auth.ImgUrl,
	)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

// SearchAuthByEmail implements repositories.AuthRepo
func (a *authRepo) SearchAuthByEmail(ctx context.Context, email string, id int) ([]*models.SearchAuth, error) {
	stmt, err := a.db.Prepare(querySearchByEmail)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id, "%"+email+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var auth []*models.SearchAuth

	for rows.Next() {
		tempAuth := new(models.SearchAuth)
		err := rows.Scan(
			&tempAuth.Email,
			&tempAuth.Id,
			&tempAuth.ImgUrl,
			&tempAuth.FollowingId,
			&tempAuth.AuthId,
		)

		if err != nil {
			return nil, err
		}

		auth = append(auth, tempAuth)
	}

	return auth, nil

}

// Create implements repositories.AuthRepo
func (a *authRepo) Create(ctx context.Context, auth *models.Auth) error {
	stmt, err := a.db.Prepare(queryCreate)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(auth.Email, auth.Password, auth.ImgUrl)
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
