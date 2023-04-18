package repository

import (
	"fmt"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(user models.User) (int, error) {
	var user_id int

	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, password) values ($1, $2, $3, $4) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.FirstName, user.LasteName, user.Email, user.Password)
	if err := row.Scan(&user_id); err != nil {
		return 0, err
	}
	return user_id, nil
}
