package postgres

import (
	"backend/apps/domain/follow/models"
	"backend/apps/domain/follow/repositories"
	"context"
	"database/sql"
	"fmt"
)

var (
	queryCreate = `
		INSERT INTO follows (authId,followingId,created_at)
		VALUES($1,$2,$3)
	`

	queryGetAllByAuthId = `
		SELECT a.email, a.id
		FROM follows as f
		JOIN auth as a
			ON a.id = f.followingId
		WHERE 
			f.authId = $1
	`

	queryUnfollow = `
		DELETE FROM follows
		WHERE 
			f.authId = $1
		AND
			f.followingId = $2
	`
)

type followRepo struct {
	db *sql.DB
}

// Create implements repositories.FollowRepo
func (f *followRepo) Create(ctx context.Context, follow *models.Follow) error {
	stmt, err := f.db.Prepare(queryCreate)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", follow)
	_, err = stmt.Exec(follow.AuthId, follow.FollowingId, follow.CreatedAt)
	return err
}

// Delete implements repositories.FollowRepo
func (f *followRepo) Delete(ctx context.Context, authId int, followingId int) error {
	stmt, err := f.db.Prepare(queryUnfollow)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(authId, followingId)
	return err
}

// GetAll implements repositories.FollowRepo
func (f *followRepo) GetAll(ctx context.Context, authId int) ([]*models.FollowWithAuth, error) {
	stmt, err := f.db.Prepare(queryGetAllByAuthId)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(authId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var follow []*models.FollowWithAuth

	for rows.Next() {
		tempFollowAuth := models.FollowWithAuth{}
		err := rows.Scan(
			&tempFollowAuth.Email,
			&tempFollowAuth.Id,
		)
		if err != nil {
			return nil, err
		}

		follow = append(follow, &tempFollowAuth)
	}

	return follow, nil
}

func NewFollowRepo(db *sql.DB) repositories.FollowRepo {
	return &followRepo{
		db: db,
	}
}
