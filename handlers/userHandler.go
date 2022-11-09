package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joeernest/gormapirest/db"
	"github.com/joeernest/gormapirest/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GetAllUsers : Getting all users
func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Database.Find(&users)
	sendData(rw, users, http.StatusOK)
}

// checkIfUserExists : "get a user by ID"
func checkIfUserExists(r *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}

	if err := db.Database.First(&user, userId); err.Error != nil {
		return user, err
	} else {
		return user, nil
	}

}

// GetUserByID : Get a user by ID
func GetUserByID(rw http.ResponseWriter, r *http.Request) {
	user, err := checkIfUserExists(r)
	if err != nil || user.Id == 0 {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)
	}
}

// CreateUser : Creates a new user
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&user)
		sendData(rw, user, http.StatusCreated)
	}
}

// UpdateUserByID : Updates a user by ID
func UpdateUserByID(rw http.ResponseWriter, r *http.Request) {
	var userId int64
	if oldUser, err := checkIfUserExists(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		userId = oldUser.Id
		user := models.User{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.Database.Save(&user)
			sendData(rw, user, http.StatusOK)
		}
	}
}

// DeleteUserByID : Delete a user by ID
func DeleteUserByID(rw http.ResponseWriter, r *http.Request) {
	if user, err := checkIfUserExists(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.Database.Delete(&user)
		sendData(rw, user, http.StatusOK)
	}
}
