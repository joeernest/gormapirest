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

// GetUser : Get a user by ID
func GetUser(rw http.ResponseWriter, r *http.Request) {
	user, err := getUserById(r)
	if err != nil || user.Id == 0 {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)
	}
}

// getUserById : "func private"
func getUserById(r *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}

	if err := db.Database.First(&user, userId); err.Error != nil {
		return user, err
	} else {
		return user, nil
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

// UpdateUser : Updates a user by ID
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var userId int64
	if oldUser, err := getUserById(r); err != nil {
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

// DeleteUser : Delete a user by ID
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.Database.Delete(&user)
		sendData(rw, user, http.StatusOK)
	}
}
