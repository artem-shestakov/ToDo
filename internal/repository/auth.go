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

	query := fmt.Sprintf("INSERT INTO %s (user_id, first_name, last_name) values ($1, $2, $3) RETURNING user_id", userTable)
	fmt.Println(user, query)
	row := r.db.QueryRow(query, user.ID, user.FirstName, user.LasteName)
	if err := row.Scan(&user_id); err != nil {
		return 0, err
	}
	return user_id, nil
}

func (r *AuthRepository) GetUser(user_id int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", userTable)
	err := r.db.Get(&user, query, user_id)
	return user, err
}
