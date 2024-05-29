package main

import (
	"database/sql"
	"fmt"
	"student-attendance/controllers"
	"student-attendance/database"
	"student-attendance/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "final"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}
	database.DbMigrate(DB)
	defer DB.Close()

	r := gin.Default()
	// Router Category
	r.GET("/user", controllers.GetAllUsers)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	// Router Class
	r.GET("/class", middleware.Auth([]string{"admin", "siswa"}), controllers.GetAllClass)
	r.POST("/class", middleware.Auth([]string{"admin"}), controllers.InsertClass)
	r.PUT("/class/:id", middleware.Auth([]string{"admin"}), controllers.UpdateClass)
	r.DELETE("/class/:id", middleware.Auth([]string{"admin"}), controllers.DeleteClass)
	// Router student
	r.GET("/student", middleware.Auth([]string{"admin", "siswa"}), controllers.GetAllStudent)
	r.POST("/student", middleware.Auth([]string{"admin"}), controllers.InsertStudent)
	r.PUT("/student/:id", middleware.Auth([]string{"admin"}), controllers.UpdateStudent)
	r.DELETE("/student/:id", middleware.Auth([]string{"admin"}), controllers.DeleteStudent)

	r.Run("localhost:8080")
}
