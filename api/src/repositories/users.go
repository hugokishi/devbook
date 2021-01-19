package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users - User Repository
type Users struct {
	db *sql.DB
}

// NewUserRepository - Create new user repository
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// CreateNewUser - Insert new user in database
func (repository Users) CreateNewUser(user models.User) (uint64, error) {

	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastIDInserted), nil
}
