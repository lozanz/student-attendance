package controllers

import (
	"net/http"
	"strconv"
	"student-attendance/database"
	"student-attendance/repository"
	"student-attendance/structs"

	"github.com/gin-gonic/gin"
)

func GetAllStudent(c *gin.Context) {
	var (
		result gin.H
	)
	student, err := repository.GetAllStudent(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": student,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertStudent(c *gin.Context) {
	var student structs.Student

	err := c.ShouldBindJSON(&student)
	if err != nil {
		panic(err)
	}
	user, err := repository.GetUserByID(database.DbConnection, student.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	student.Name = user.Username

	if err := repository.InsertStudent(database.DbConnection, student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert Student"})
}

func UpdateStudent(c *gin.Context) {
	var student structs.Student

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&student)
	if err != nil {
		panic(err)
	}

	student.ID = int64(id)

	err = repository.UpdateStudent(database.DbConnection, student)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Student",
	})
}

func DeleteStudent(c *gin.Context) {
	var student structs.Student
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	student.ID = int64(id)

	err = repository.DeleteStudent(database.DbConnection, student)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Student",
	})
}
