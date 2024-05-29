package controllers

import (
	"net/http"
	"strconv"
	"student-attendance/database"
	"student-attendance/repository"
	"student-attendance/structs"

	"github.com/gin-gonic/gin"
)

func GetAllClass(c *gin.Context) {
	var (
		result gin.H
	)
	class, err := repository.GetAllClass(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": class,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertClass(c *gin.Context) {
	var class structs.Class

	err := c.ShouldBindJSON(&class)
	if err != nil {
		panic(err)
	}

	err1, Class := repository.GetAllClass(database.DbConnection)
	if err1 != nil {
		panic(err1)
	}

	class.ID = 0
	for _, c := range Class {
		if c.ID > class.ID {
			class.ID = c.ID
		}
	}
	class.ID++

	err = repository.InsertClass(database.DbConnection, class)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Class",
	})
}

func UpdateClass(c *gin.Context) {
	var class structs.Class

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&class)
	if err != nil {
		panic(err)
	}

	class.ID = int64(id)

	err = repository.UpdateClass(database.DbConnection, class)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Class",
	})
}

func DeleteClass(c *gin.Context) {
	var class structs.Class
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	class.ID = int64(id)

	err = repository.DeleteClass(database.DbConnection, class)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Class",
	})
}
