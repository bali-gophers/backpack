package main

import (
	"database/sql"
)

type ProfileRepo struct {
	sqlDB *sql.DB
}

const (
	selectCountQuery = `
		SELECT COUNT(*) FROM github_profile
	`
	selectQuery = `
		SELECT * FROM github_profile
	`

	insertQuery = `
		INSERT INTO github_profile (
			id, name, email, avatar_url, bio, public_repos, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	updateQuery = `
		UPDATE github_profile SET
			name = ?,
			email = ?,
			avatar_url = ?,
			bio = ?,
			public_repos = ?,
			created_at = ?
		WHERE id = ?
	`
)

func NewProfileRepo(sqlDB *sql.DB) ProfileRepo {
	return ProfileRepo{
		sqlDB: sqlDB,
	}
}

func (repo ProfileRepo) Store(profile Profile) error {
	exists, err := repo.existsByID(profile.ID)
	if err != nil {
		return err
	}
	if !exists {
		if err := repo.insert(profile); err != nil {
			return err
		}
	}
	if err := repo.update(profile); err != nil {
		return err
	}
	return nil
}

func (repo ProfileRepo) existsByID(profileID int64) (bool, error) {
	var total int
	row := repo.sqlDB.QueryRow(selectCountQuery+" WHERE id = ?", profileID)
	if err := row.Scan(&total); err != nil {
		return false, err
	}
	return total > 0, nil
}

func (repo ProfileRepo) insert(profile Profile) error {
	stmt, err := repo.sqlDB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(
		profile.ID,
		profile.Name,
		profile.Email,
		profile.AvatarURL,
		profile.Bio,
		profile.PublicRepos,
		profile.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (repo ProfileRepo) update(profile Profile) error {
	stmt, err := repo.sqlDB.Prepare(updateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(
		profile.Name,
		profile.Email,
		profile.AvatarURL,
		profile.Bio,
		profile.PublicRepos,
		profile.CreatedAt,
		profile.ID); err != nil {
		return err
	}
	return nil
}

func (repo ProfileRepo) ResolveByID(profileID int64) (Profile, error) {
	var prof Profile
	row := repo.sqlDB.QueryRow(selectQuery+" WHERE id = ?", profileID)
	if err := row.Scan(
		&prof.ID,
		&prof.Name,
		&prof.Email,
		&prof.AvatarURL,
		&prof.Bio,
		&prof.PublicRepos,
		&prof.CreatedAt); err != nil {
		return prof, err
	}
	return prof, nil
}
