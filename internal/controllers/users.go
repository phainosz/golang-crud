package controllers

import (
	"context"
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
	repository := repositories.New(repositories.ConnectionOption{ConnectionSql: db})

	if err = repository.User.CreateUser(context.Background(), user); err != nil {
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
	repository := repositories.New(repositories.ConnectionOption{ConnectionSql: db})
	users, err := repository.User.GetUsers(context.Background())
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

	repository := repositories.New(repositories.ConnectionOption{ConnectionSql: db})
	if err = repository.User.UpdateUser(context.Background(), id, user); err != nil {
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

	repository := repositories.New(repositories.ConnectionOption{ConnectionSql: db})
	if err = repository.User.DeleteUserById(context.Background(), id); err != nil {
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

	repository := repositories.New(repositories.ConnectionOption{ConnectionSql: db})
	user, err := repository.User.FindUserById(context.Background(), id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, utils.ErrorResponse{Error: err.Error()})
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, user)
}
