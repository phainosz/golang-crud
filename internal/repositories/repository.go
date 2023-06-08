package repositories

import (
	"database/sql"

	"github.com/phainosz/golang-crud/internal/repositories/users"
)

type ConnectionOption struct {
	ConnectionSql *sql.DB
}

type Container struct {
	User users.UserRepository
	//Job jobs.JobRepository to add new repositories
}

func New(conn ConnectionOption) *Container {
	return &Container{
		User: users.NewUserRepository(conn.ConnectionSql),
		//Job jobs.NewJobRepository to add new repositories
	}
}
