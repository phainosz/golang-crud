package repositories

import (
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

	lines, err := db.Query("select * from users")

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
			&user.Email,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

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

func DeleteUserById(ID uint64) error {
	db, err := db.Connect()

	if err != nil {
		log.Fatal(err)
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
