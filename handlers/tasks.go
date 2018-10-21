package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
	"github.com/mrgleam/sec-unit-tests-example/models"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenInCookie, err := c.Cookie("token")
		if err != nil {
			return err
		}
		userInToken := models.Token{}
		// Let's parse this by the secrete, which only server knows.
		_, err = jwt.ParseWithClaims(tokenInCookie.Value, &userInToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			log.Println(err)
			return err
		}
		userID := models.GetUserID(db, userInToken.Email)
		// Fetch tasks using our new model
		return c.JSON(http.StatusOK, models.GetTasks(db,userID))
	}
}

// PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var task models.Task
		// Map imcoming JSON body to the new Task
		c.Bind(&task)
		tokenInCookie, err := c.Cookie("token")
		if err != nil {
			return err
		}
		userInToken := models.Token{}
		// Let's parse this by the secrete, which only server knows.
		_, err = jwt.ParseWithClaims(tokenInCookie.Value, &userInToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			log.Println(err)
			return err
		}
		userID := models.GetUserID(db, userInToken.Email)
		// Add a task using our new model
		id, err := models.PutTask(db, task.Name, userID)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		_, err := models.DeleteTask(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// Handle errors
		} else {
			return err
		}
	}
}
