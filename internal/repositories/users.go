package repositories

import (
	"errors"
	"log"

	"github.com/phainosz/golang-crud/internal/db"
	"github.com/phainosz/golang-crud/internal/models"
)

// get all users from database
func GetUsers() ([]models.User, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
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
func CreateUser(user models.User) {
	db, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		log.Fatal("Error preparing statement")
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email)
	if err != nil {
		log.Fatal("Error executing statement")
	}
}

// delete user by id
func DeleteUserById(ID uint64) error {
	db, err := db.Connect()

	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

// update user by ID
func UpdateUser(ID uint64, user models.User) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// find user by id
func FindUserById(ID uint64) (models.User, error) {
	db, err := db.Connect()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	rows, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {

		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return models.User{}, err
		}
	}

	if user.ID == 0 {
		return user, errors.New("ID not found")
	}

	return user, nil
}
