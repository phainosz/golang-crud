package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/phainosz/golang-crud/internal/db"
	"github.com/phainosz/golang-crud/internal/models"
	"github.com/phainosz/golang-crud/internal/repositories"
	"github.com/phainosz/golang-crud/internal/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnprocessableEntity, utils.ErrorResponse{Error: err.Error()})
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	db, err := db.Connect()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
	}
	defer db.Close()
	userRepository := repositories.NewUserRepository(db)

	if err = userRepository.CreateUser(user); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}

	utils.WriteSuccessResponse(w, http.StatusCreated, nil)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	defer db.Close()
	userRepository := repositories.NewUserRepository(db)
	users, err := userRepository.GetUsers()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}

	if len(users) <= 0 {
		utils.WriteErrorResponse(w, http.StatusNotFound, utils.ErrorResponse{Error: "users not found"})
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnprocessableEntity, utils.ErrorResponse{Error: err.Error()})
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	db, err := db.Connect()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	if err = userRepository.UpdateUser(id, user); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	utils.WriteSuccessResponse(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	db, err := db.Connect()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	if err = userRepository.DeleteUserById(id); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	utils.WriteSuccessResponse(w, http.StatusNoContent, nil)
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	db, err := db.Connect()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	user, err := userRepository.FindUserById(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, utils.ErrorResponse{Error: err.Error()})
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, user)
}
