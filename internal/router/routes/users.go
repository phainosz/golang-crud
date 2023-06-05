package routes

import (
	"net/http"

	"github.com/phainosz/golang-crud/internal/controllers"
)

var usersRoutes = []Route{
	{
		Uri:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
	},
	{
		Uri:      "/users",
		Method:   http.MethodGet,
		Function: controllers.GetAllUsers,
	},
	{
		Uri:      "/users/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
	},
	{
		Uri:      "/users/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
	},
	{
		Uri:      "/users/{id}",
		Method:   http.MethodGet,
		Function: controllers.FindUserById,
	},
}
