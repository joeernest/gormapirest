package models

import (
	"fmt"
	"github.com/joeernest/gormapirest/db"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

// MigrateUsers : Auto-migrate the table User
func MigrateUsers() {
	if err := db.Database.AutoMigrate(User{}); err != nil {
		fmt.Printf("Error migrating users table: %v", err)
	}
}
