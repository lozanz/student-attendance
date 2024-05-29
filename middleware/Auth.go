package middleware

import (
	"net/http"
	"student-attendance/database"
	"student-attendance/repository"
	"student-attendance/structs"

	"github.com/gin-gonic/gin"
)

func Auth(allowedRole []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resultString string
		responseHeader := http.StatusBadRequest

		username, password, ok := c.Request.BasicAuth()

		if !ok {
			resultString = "username and password is required"
		} else {
			user := structs.Users{
				Username: username,
				Password: password,
			}

			userRowData, err := repository.AuthenticateUser(database.DbConnection, user)

			if err != nil {
				resultString = err.Error()
			} else {
				if len(userRowData) > 0 {
					isAuthorized := false

					for _, role := range allowedRole {
						if role == "admin" || role == "siswa" {
							if role == userRowData[0].Role {
								isAuthorized = true
								break
							}
						}
					}

					if isAuthorized || len(allowedRole) == 0 {
						c.Next()
						return
					} else {
						resultString = "you're not allowed to access this route"
						responseHeader = http.StatusUnauthorized
					}
				} else {
					resultString = "username and password is incorrect"
				}
			}
		}

		c.JSON(responseHeader, gin.H{
			"result": resultString,
		})
		c.Abort()
	}
}
