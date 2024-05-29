package controllers

import (
	"net/http"
	"strconv"
	"student-attendance/database"
	"student-attendance/repository"
	"student-attendance/structs"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllAttendance(c *gin.Context) {
	var (
		result gin.H
	)
	attendance, err := repository.GetAllAttendance(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": attendance,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertAttendance(c *gin.Context) {
	var attendance structs.Attendance

	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validStatuses := map[string]bool{"hadir": true, "absen": true, "terlambat": true, "izin": true, "sakit": true}
	if !validStatuses[attendance.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	_, err := time.Parse("2006-01-02", attendance.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	if err := repository.InsertAttendance(database.DbConnection, attendance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert Attendance"})
}

func UpdateAttendance(c *gin.Context) {
	var attendance structs.Attendance
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attendance.ID = int64(id)
	validStatuses := map[string]bool{"hadir": true, "absen": true, "terlambat": true, "izin": true, "sakit": true}
	if !validStatuses[attendance.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	_, err := time.Parse("2006-01-02", attendance.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	if err := repository.UpdateAttendance(database.DbConnection, attendance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Update Attendance"})
}

func DeleteAttendance(c *gin.Context) {
	var attendance structs.Attendance
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	attendance.ID = int64(id)

	err = repository.DeleteAttendance(database.DbConnection, attendance)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Attendance",
	})
}
