package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// GetUsers - Returns users who meet the nick or name filter
func (repository Users) GetUsers(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByID - Get one user based on ID
func (repository Users) GetByID(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// UpdateUser - Method to update user in database
func (repository Users) UpdateUser(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// DeleteUser - Delete user in database
func (repository Users) DeleteUser(ID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from users where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// GetByEmail - Get user in database based on email
func (repository Users) GetByEmail(email string) (models.User, error) {
	lines, err := repository.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, nil
		}
	}

	return user, nil
}

// FollowUser - Method to follow other user
func (repository Users) FollowUser(userID, followID uint64) error {
	statement, err := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followID); err != nil {
		return err
	}

	return nil
}

// UnFollowUser - Method to unfollow user
func (repository Users) UnFollowUser(userID, followID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followID); err != nil {
		return err
	}

	return nil
}
