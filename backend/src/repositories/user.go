package repositories

import (
	"database/sql"
	"errors"
	"shared-blog-backend/src/models"

	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	db *sql.DB
}

func UserRepository(database *sql.DB) *UserDB {
	return &UserDB{database}
}

func (userDb *UserDB) SignUpUser(user models.User) error {

	statement, err := userDb.db.Prepare(`INSERT INTO users (name, email, username, password) VALUES (?,?,?,?) `)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, user.Username, user.Password); err != nil {
		return err
	}
	return nil
}

func (userDb *UserDB) SignInUser(userSent models.User) (uint64, error) {
	var user models.User

	row, err := userDb.db.Query("SELECT id, name, email, username, password FROM users WHERE email = ?", userSent.Email)

	if err != nil {
		return 0, err
	}

	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Password); err != nil {
			return 0, err
		}
	}

	if user.ID == 0 {
		return 0, errors.New("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userSent.Password)); err != nil {
		return 0, errors.New("incorrect password")
	}

	return user.ID, nil
}
