package controllers

import (
	"net/http"
	"student-attendance/database"
	"student-attendance/repository"
	"student-attendance/structs"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var (
		result gin.H
	)
	user, err := repository.GetAllUsers(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": user,
		}
	}
	c.JSON(http.StatusOK, result)
}

func RegisterUser(c *gin.Context) {
	var user structs.Users
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	if user.Role != "admin" && user.Role != "siswa" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role must 'admin' or 'siswa'"})
		return
	}

	err1, users := repository.GetAllUsers(database.DbConnection)
	if err1 != nil {
		panic(err1)
	}

	user.ID = 0
	for _, c := range users {
		if c.ID > user.ID {
			user.ID = c.ID
		}
	}
	user.ID++

	if err := repository.RegisterUser(database.DbConnection, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
	var loginUser structs.Users
	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		panic(err)
	}

	user, err := repository.LoginUser(database.DbConnection, loginUser)
	if err != nil {
		panic(err)
	}

	if user.Username == "" || user.Password != loginUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
