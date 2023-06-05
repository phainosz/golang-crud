package repositories

import (
	"database/sql"
	"errors"

	"github.com/phainosz/golang-crud/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// get all users from database
func (userRepository UserRepository) GetUsers() ([]models.User, error) {
	rows, err := userRepository.db.Query("select * from users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

// create user
func (userRepository UserRepository) CreateUser(user models.User) error {
	statement, err := userRepository.db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

// delete user by id
func (userRepository UserRepository) DeleteUserById(id uint64) error {
	statement, err := userRepository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}
	return nil
}

// update user by id
func (userRepository UserRepository) UpdateUser(id uint64, user models.User) error {
	statement, err := userRepository.db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Email, id); err != nil {
		return err
	}

	return nil
}

// find user by id
func (userRepository UserRepository) FindUserById(id uint64) (models.User, error) {
	rows, err := userRepository.db.Query("select * from users where id = ?", id)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {

		if err = rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return models.User{}, err
		}
	}

	if user.Id == 0 {
		return user, errors.New("id not found")
	}

	return user, nil
}
