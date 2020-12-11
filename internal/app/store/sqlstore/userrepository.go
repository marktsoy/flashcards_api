package sqlstore

import (
	"database/sql"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *SQLStore
}

// Create ...
func (r *UserRepository) Create(u *models.User) error {
	db := r.store.db

	if err := u.Creating(); err != nil {
		return err
	}
	err := db.QueryRow(
		"INSERT INTO users(email,encrypted_password) VALUES ($1,$2) RETURNING id",
		u.Email, u.EncryptedPassword,
	).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id,email,encrypted_password FROM users WHERE email=$1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	users := []*models.User{}

	rows, err := r.store.db.Query("SELECT id,email FROM users")
	if err != nil {
		return nil, store.ErrCanNotGetResults
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, store.ErrCanNotGetResults
		}
		users = append(users, user)
	}

	return users, nil
}
