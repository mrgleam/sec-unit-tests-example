package models

import (
	"database/sql"
	"log"
)

// User is a struct containing User data
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUser(db *sql.DB, email string) User {
	result := User{}
	sqlStatement := `SELECT * FROM users WHERE email=$1`
	row := db.QueryRow(sqlStatement, email)
	err := row.Scan(&result.ID, &result.Email, &result.Password)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Zero rows found")
		} else {
			panic(err)
		}
	}
	return result
}

func GetUserID(db *sql.DB, email string) int {
	var id int
	sqlStatement := `SELECT id FROM users WHERE email=$1`
	row := db.QueryRow(sqlStatement, email)
	err := row.Scan(&id)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Zero rows found")
		} else {
			panic(err)
		}
	}
	return id
}
