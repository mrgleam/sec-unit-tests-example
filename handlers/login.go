package handlers

import (
	"database/sql"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mrgleam/sec-unit-tests-example/models"
)

// Login endpoint
func Login(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		c.Bind(&user)
		userInDB := models.GetUser(db, user.Email)
		if user.Email == userInDB.Email && models.ComparePasswords(userInDB.Password, []byte(user.Password)) {
			// Create token
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["email"] = userInDB.Email
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}

		return echo.ErrUnauthorized
	}
}

// func Login(c echo.Context) error {
// 	var user models.User
// 	c.Bind(&user)
// 	GetUser(db, )
// 	if user.Email == "jon" && user.Password == "shhh!" {
// 		// Create token
// 		token := jwt.New(jwt.SigningMethodHS256)

// 		// Set claims
// 		claims := token.Claims.(jwt.MapClaims)
// 		claims["name"] = "Jon Snow"
// 		claims["admin"] = true
// 		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

// 		// Generate encoded token and send it as response.
// 		t, err := token.SignedString([]byte("secret"))
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"token": t,
// 		})
// 	}

// 	return echo.ErrUnauthorized
// }
