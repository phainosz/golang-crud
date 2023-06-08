package users

import (
	"context"
	"database/sql"
	"errors"

	"github.com/phainosz/golang-crud/internal/models"
)

type userSql struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userSql{db}
}

// get all users from database
func (repo userSql) GetUsers(ctx context.Context) ([]models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "select * from users")

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
func (repo userSql) CreateUser(ctx context.Context, user models.User) error {
	statement, err := repo.db.PrepareContext(ctx, "insert into users (name, email) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.ExecContext(ctx, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

// delete user by id
func (repo userSql) DeleteUserById(ctx context.Context, id uint64) error {
	statement, err := repo.db.PrepareContext(ctx, "delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.ExecContext(ctx, id); err != nil {
		return err
	}
	return nil
}

// update user by id
func (repo userSql) UpdateUser(ctx context.Context, id uint64, user models.User) error {
	statement, err := repo.db.PrepareContext(ctx, "update users set name = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.ExecContext(ctx, user.Name, user.Email, id); err != nil {
		return err
	}

	return nil
}

// find user by id
func (repo userSql) FindUserById(ctx context.Context, id uint64) (models.User, error) {
	rows := repo.db.QueryRowContext(ctx, "select * from users where id = ?", id)
	if err := rows.Err(); err != nil {
		return models.User{}, err
	}

	var user models.User

	if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("id not found")
		}
		return models.User{}, err
	}

	return user, nil
}
